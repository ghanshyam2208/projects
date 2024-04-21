import { Module } from '@nestjs/common';
import { SongsController } from './songs.controller';
import { SongsService } from './songs.service';
import { CommonModule } from 'src/common';
import { SongsRepository } from './songs.repository';

@Module({
  imports: [CommonModule],
  controllers: [SongsController],
  providers: [SongsService, SongsRepository],
})
export class SongsModule {}
