import {
  Body,
  Controller,
  DefaultValuePipe,
  Delete,
  Get,
  Param,
  ParseIntPipe,
  Post,
  Put,
  Query,
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

  @Get('')
  findAll(
    @Query('page', new DefaultValuePipe(1), ParseIntPipe) page: number,
    @Query('limit', new DefaultValuePipe(2), ParseIntPipe) limit: number,
  ) {
    return this.songsService.findAll(page, limit);
  }

  @Get(':id')
  findOne(@Param('id', ParseIntPipe) id: number) {
    return this.songsService.findOne(id);
  }

  @Put(':id')
  updateOne(
    @Param('id', ParseIntPipe) id: number,
    @Body(new CreateSongValidationPipe()) updateSongPayload: CreateSongPayload,
  ) {
    return this.songsService.updateOne(id, updateSongPayload);
  }

  @Delete(':id')
  deleteOne(@Param('id', ParseIntPipe) id: number) {
    return this.songsService.deleteOne(id);
  }
}
