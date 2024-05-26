import { ConfigModule, ConfigService } from '@nestjs/config';
import { join } from 'path';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { TODO_PACKAGE_NAME } from 'proto/todo';
import { AUTH_PACKAGE_NAME } from 'proto/auth';

export const ClientsModuleImports = ClientsModule.registerAsync([
  {
    imports: [ConfigModule],
    name: 'MULTI_PACKAGE_LOOKUP_NAME',
    useFactory: async (configService: ConfigService) => ({
      transport: Transport.GRPC,
      options: {
        protoPath: [
          join(__dirname, '../proto/todo.proto'),
          join(__dirname, '../proto/auth.proto'),
        ],
        url: `${configService.get('GRPC_HOST_IP')}:${configService.get('GRPC_PORT')}`,
        package: [TODO_PACKAGE_NAME, AUTH_PACKAGE_NAME],
      },
    }),
    inject: [ConfigService],
  },
  {
    imports: [ConfigModule],
    name: 'NOTIFICATION_SERVICE',
    useFactory: async (configService: ConfigService) => ({
      transport: Transport.RMQ,
      options: {
        urls: [
          `amqp://${configService.get<string>('NOTIFICATION_RMQ_HOST_IP')}:${configService.get<string>('NOTIFICATION_RMQ_HOST_PORT')}`,
        ],
        queue: 'notifications',
      },
    }),
    inject: [ConfigService],
  },
]);

export const userOtpGenerator = (() =>
  Math.floor(Math.random() * 900000) + 100000)().toString();
