import { Injectable, NotFoundException } from '@nestjs/common';
import { v4 as uuidv4 } from 'uuid';
import { CreateSongPayload } from './songs.validations';

@Injectable()
export class SongsService {
  private readonly songs = {};

  createSong(createSongPayload: CreateSongPayload) {
    const id = uuidv4();
    const song = {
      ...createSongPayload,
      id,
    };
    this.songs[id] = song;
    return song;
  }

  findAll() {
    return Object.keys(this.songs).map((key) => {
      return this.songs[key];
    });
  }

  findOne(id: string) {
    const song = this.songs[id];
    if (!song) {
      throw new NotFoundException(`no song found with id ${id}`);
    }
    return song;
  }

  updateOne(id: string, updateSongPayload: CreateSongPayload) {
    const song = this.songs[id];
    if (!song) {
      throw new NotFoundException(`no song found with id ${id}`);
    }
    this.songs[id] = {
      ...updateSongPayload,
      id,
    };
    return this.songs[id];
  }

  deleteOne(id: string) {
    const song = this.songs[id];
    if (!song) {
      throw new NotFoundException(`no song found with id ${id}`);
    }
    delete this.songs[id];
    return song;
  }
}
