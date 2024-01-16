# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import weather_pb2 as weather__pb2


class WeatherServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StreamForecasts = channel.unary_stream(
                '/weather.WeatherService/StreamForecasts',
                request_serializer=weather__pb2.ForecastRequest.SerializeToString,
                response_deserializer=weather__pb2.ForecastResponse.FromString,
                )


class WeatherServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def StreamForecasts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_WeatherServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StreamForecasts': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamForecasts,
                    request_deserializer=weather__pb2.ForecastRequest.FromString,
                    response_serializer=weather__pb2.ForecastResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'weather.WeatherService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class WeatherService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def StreamForecasts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/weather.WeatherService/StreamForecasts',
            weather__pb2.ForecastRequest.SerializeToString,
            weather__pb2.ForecastResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)