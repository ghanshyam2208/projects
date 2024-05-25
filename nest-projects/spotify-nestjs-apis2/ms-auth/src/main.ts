import { NestFactory } from '@nestjs/core';
import { AuthModule } from './auth.module';
// import { ConfigService } from '@nestjs/config';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AuthModule,
    {
      transport: Transport.GRPC,
      options: {
        protoPath: join(__dirname, '../proto/todo.proto'),
        url: '0.0.0.0:22051',
        package: 'todo',
      },
    },
  );
  // const configService = app.get(ConfigService);
  // await app.listen(configService.get('HTTP_PORT'));
  app.listen();
}
bootstrap();
