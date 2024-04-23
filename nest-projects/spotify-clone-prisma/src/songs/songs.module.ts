import { Module } from '@nestjs/common';
import { SongsController } from './songs.controller';
import { SongsService } from './songs.service';
import { CommonModule } from 'src/common';
import { SongsRepository } from './songs.repository';
import { AuthModule } from 'src/auth/auth.module';
import { JwtService } from '@nestjs/jwt';

@Module({
  imports: [CommonModule, AuthModule],
  controllers: [SongsController],
  providers: [SongsService, SongsRepository, JwtService],
})
export class SongsModule {}
