import { Module } from '@nestjs/common';
import { SongsController } from './songs.controller';
import { SongsService } from './songs.service';
import { SongsRepository } from './songs.repo';
import { PrismaService } from 'src/common';

@Module({
  imports: [],
  controllers: [SongsController],
  providers: [SongsService, SongsRepository, PrismaService],
})
export class SongsModule {}
