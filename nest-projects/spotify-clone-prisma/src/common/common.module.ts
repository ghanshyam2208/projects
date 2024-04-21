import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { PrismaService } from './prisma.service';
import { CryptoHelper } from './crypto.helper';

@Module({
  imports: [ConfigModule.forRoot()],
  providers: [PrismaService, CryptoHelper],
  exports: [PrismaService, CryptoHelper],
})
export class CommonModule {}
