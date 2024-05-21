import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { Transport } from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.connectMicroservice({
    transport: Transport.RMQ,
    options: {
      urls: ['amqp://0.0.0.0:5672'],
      queue: 'analytics',
    },
  });
  await app.startAllMicroservices();
  await app.listen(3003);
}
bootstrap();
