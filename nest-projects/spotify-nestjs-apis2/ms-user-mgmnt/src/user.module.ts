import { Module } from '@nestjs/common';
import { UserController } from './user.controller';
import { UserService } from './user.service';
import { ConfigModule } from '@nestjs/config';
import { PrismaModule } from './prisma/prisma.module';
import { UserRepository } from './user.repository';
import { CryptoHelper } from './prisma/crypto.helper';
import { CommonModule } from './common/common.module';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ConfigModule,
    PrismaModule,
    CommonModule,
    ClientsModule.register([
      {
        name: 'todo',
        transport: Transport.GRPC,
        options: {
          protoPath: join(__dirname, '../proto/todo.proto'),
          url: '0.0.0.0:22051',
          package: 'todo',
        },
      },
    ]),
  ],
  controllers: [UserController],
  providers: [UserService, UserRepository, CryptoHelper],
})
export class UserModule {}
