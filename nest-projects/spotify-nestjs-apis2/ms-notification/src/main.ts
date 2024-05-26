import { NestFactory } from '@nestjs/core';
import { NotificationModule } from './notification.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  // Create a new NestJS application instance
  const app = await NestFactory.create(NotificationModule);

  // Get an instance of the ConfigService from the application
  const configService = app.get(ConfigService);

  const amqpUrl = `amqp://${configService.get<string>('NOTIFICATION_RMQ_HOST_IP')}:${configService.get<string>('NOTIFICATION_RMQ_HOST_PORT')}`;

  // Connect the application to the RMQ microservice
  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.RMQ,
    options: {
      urls: [amqpUrl],
      queue: 'notifications',
    },
  });

  // Start all microservices
  await app.startAllMicroservices();
}
bootstrap();
