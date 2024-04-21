import { Body, Controller, Post } from '@nestjs/common';
import { SongsService } from './songs.service';
import {
  CreateSongPayload,
  CreateSongValidationPipe,
} from './songs.validations';

@Controller('songs')
export class SongsController {
  constructor(private readonly songsService: SongsService) {}

  @Post()
  createSong(
    @Body(new CreateSongValidationPipe()) createPostPayload: CreateSongPayload,
  ) {
    return this.songsService.createSong(createPostPayload);
  }
}
