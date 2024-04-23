import { Body, Controller, Post, UseGuards } from '@nestjs/common';
import { SongsService } from './songs.service';
import {
  CreateSongPayload,
  CreateSongValidationPipe,
} from './songs.validations';
import { AuthGuard } from 'src/auth/auth.guard';

@Controller('songs')
export class SongsController {
  constructor(private readonly songsService: SongsService) {}

  @UseGuards(AuthGuard)
  @Post()
  createSong(
    @Body(new CreateSongValidationPipe()) createPostPayload: CreateSongPayload,
  ) {
    return this.songsService.createSong(createPostPayload);
  }
}
