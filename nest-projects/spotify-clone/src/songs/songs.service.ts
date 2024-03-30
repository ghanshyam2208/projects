import { Inject, Injectable } from '@nestjs/common';

import { CreateSongPayload } from './songs.validations';
import { MySqlConnectionType } from 'src/common/constant/mysql.connection';
import { DeleteResult, Repository, UpdateResult } from 'typeorm';
import { SongsModel } from './songs.model';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class SongsService {
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

  findAll(): Promise<SongsModel[]> {
    return this.songsRepository.find({});
  }

  findOne(id: number): Promise<SongsModel> {
    return this.songsRepository.findOne({
      where: {
        id,
      },
    });
  }

  updateOne(
    id: number,
    updateSongPayload: CreateSongPayload,
  ): Promise<UpdateResult> {
    return this.songsRepository.update(id, updateSongPayload);
  }

  deleteOne(id: number): Promise<DeleteResult> {
    return this.songsRepository.delete({
      id,
    });
  }
}
