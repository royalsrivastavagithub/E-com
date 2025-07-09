from fastapi import FastAPI
from pydantic import BaseModel
import grpc

from protobuf import hello_pb2, hello_pb2_grpc

app = FastAPI()


# Define the request schema for Swagger docs
class HelloInput(BaseModel):
    name: str


@app.post("/")
async def say_hello(input: HelloInput):
    # gRPC channel to the gRPC server
    async with grpc.aio.insecure_channel("localhost:50051") as channel:
        stub = hello_pb2_grpc.HelloServiceStub(channel)
        request = hello_pb2.HelloRequest(name=input.name)
        response = await stub.SayHello(request)

    return {"response": response.data}
