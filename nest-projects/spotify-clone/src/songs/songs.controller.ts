import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Put,
} from '@nestjs/common';
import { CreateSongValidationPipe } from './songs.validations';

@Controller('songs')
export class SongsController {
  @Post()
  createSong(@Body(new CreateSongValidationPipe()) createPostPayload: any) {
    return `returns all songs ${createPostPayload}`;
  }

  @Get()
  findAll() {
    return 'returns all songs';
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return `returns a songs ${id}`;
  }

  @Put(':id')
  updateOne(@Param('id') id: string) {
    return `updates song ${id}`;
  }

  @Delete(':id')
  deleteOne(@Param('id') id: string) {
    return `deletes song ${id}`;
  }
}
