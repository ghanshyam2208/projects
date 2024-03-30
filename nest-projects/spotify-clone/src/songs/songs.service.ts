import { Inject, Injectable, NotFoundException } from '@nestjs/common';

import { CreateSongPayload } from './songs.validations';
import { MySqlConnectionType } from 'src/common/constant/mysql.connection';
import { Repository } from 'typeorm';
import { SongsModel } from './songs.model';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class SongsService {
  private readonly songs = {};

  constructor(
    @Inject('MYSQL_CONNECTION') mySqlConnection: MySqlConnectionType,
    @InjectRepository(SongsModel)
    private songsRepository: Repository<SongsModel>,
  ) {
    console.log(mySqlConnection);
  }

  createSong(createSongPayload: CreateSongPayload): Promise<SongsModel> {
    const newSong = new SongsModel();
    newSong.title = createSongPayload.title;
    newSong.artists = createSongPayload.artists;
    newSong.releasedDate = createSongPayload.releasedDate;
    newSong.duration = createSongPayload.duration;
    newSong.lyrics = createSongPayload.lyrics;

    return this.songsRepository.save(newSong);
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
