import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Put,
} from '@nestjs/common';
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

  @Get()
  findAll() {
    return this.songsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.songsService.findOne(id);
  }

  @Put(':id')
  updateOne(
    @Param('id') id: string,
    @Body(new CreateSongValidationPipe()) updateSongPayload: CreateSongPayload,
  ) {
    return this.songsService.updateOne(id, updateSongPayload);
  }

  @Delete(':id')
  deleteOne(@Param('id') id: string) {
    return this.songsService.deleteOne(id);
  }
}
