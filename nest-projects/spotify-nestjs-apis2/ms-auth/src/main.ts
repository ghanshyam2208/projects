import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

// Import the ConfigService to retrieve configuration values
import { ConfigService } from '@nestjs/config';
import { AuthModule } from './auth.module';
import { TODO_PACKAGE_NAME } from 'proto/todo';
async function bootstrap() {
  // Create a new NestJS application instance
  const app = await NestFactory.create(AuthModule);

  // Get an instance of the ConfigService from the application
  const configService = app.get(ConfigService);

  // Retrieve the GRPC host IP and port from the configuration
  const grpcUrl = `${configService.get('GRPC_HOST_IP')}:${configService.get('GRPC_PORT')}`;

  // Connect the application to the GRPC microservice
  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.GRPC,
    options: {
      // Set the proto file path
      protoPath: join(__dirname, '../proto/todo.proto'),
      // Set the URL for the GRPC microservice
      url: grpcUrl,
      // Set the package name for the GRPC service
      package: TODO_PACKAGE_NAME,
    },
  });

  // Start all microservices
  await app.startAllMicroservices();
}

// Call the bootstrap function to start the application
bootstrap();
