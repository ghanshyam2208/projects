import { Module } from '@nestjs/common';
import { UserController } from './user.controller';
import { UserService } from './user.service';
import { ConfigModule } from '@nestjs/config';
import { PrismaModule } from './prisma/prisma.module';
import { UserRepository } from './user.repository';
import { CryptoHelper } from './prisma/crypto.helper';
import { CommonModule } from './common/common.module';
import { ClientsModuleImports } from './user.helpers';

@Module({
  imports: [ConfigModule, PrismaModule, CommonModule, ClientsModuleImports],
  controllers: [UserController],
  providers: [UserService, UserRepository, CryptoHelper],
})
export class UserModule {}
