# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: weather.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rweather.proto\x12\x07weather\"-\n\x0f\x46orecastRequest\x12\x0c\n\x04\x63ity\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ys\x18\x02 \x01(\x05\"M\n\x10\x46orecastResponse\x12\x0c\n\x04\x64\x61te\x18\x01 \x01(\t\x12\x0f\n\x07summary\x18\x02 \x01(\t\x12\x1a\n\x12temperatureCelsius\x18\x03 \x01(\x02\x32Z\n\x0eWeatherService\x12H\n\x0fStreamForecasts\x12\x18.weather.ForecastRequest\x1a\x19.weather.ForecastResponse0\x01\x42\x13Z\x11weather/weatherpbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'weather_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\021weather/weatherpb'
  _globals['_FORECASTREQUEST']._serialized_start=26
  _globals['_FORECASTREQUEST']._serialized_end=71
  _globals['_FORECASTRESPONSE']._serialized_start=73
  _globals['_FORECASTRESPONSE']._serialized_end=150
  _globals['_WEATHERSERVICE']._serialized_start=152
  _globals['_WEATHERSERVICE']._serialized_end=242
# @@protoc_insertion_point(module_scope)
