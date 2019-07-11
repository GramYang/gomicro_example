package main

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	tracer "gomicro_example/part7/plugins/tracer/jaeger"
	"gomicro_example/part7/plugins/tracer/opentracing/stdhttp"
	"log"
)

func init() {
	_ = plugin.Register(cors.NewPlugin())

	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
}

const name = "API gateway"

func main() {
	stdhttp.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	_ = cmd.Init()
}