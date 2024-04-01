import { Body, Controller, Post } from '@nestjs/common';
import {
  CreateSongPayload,
  CreateSongValidationPipe,
} from './songs.validations';
import { SongsService } from './songs.service';

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
