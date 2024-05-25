import { ConfigModule, ConfigService } from '@nestjs/config';
import { join } from 'path';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { TODO_PACKAGE_NAME } from 'proto/todo';

export const ClientsModuleImports = ClientsModule.registerAsync([
  {
    imports: [ConfigModule],
    name: TODO_PACKAGE_NAME,
    useFactory: async (configService: ConfigService) => ({
      transport: Transport.GRPC,
      options: {
        protoPath: join(__dirname, '../proto/todo.proto'),
        url: `${configService.get('GRPC_HOST_IP')}:${configService.get('GRPC_PORT')}`,
        package: TODO_PACKAGE_NAME,
      },
    }),
    inject: [ConfigService],
  },
]);
