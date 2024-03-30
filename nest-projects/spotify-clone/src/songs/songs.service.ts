import { Inject, Injectable } from '@nestjs/common';

import { CreateSongPayload } from './songs.validations';
import { MySqlConnectionType } from 'src/common/constant/mysql.connection';
import { DeleteResult, Repository, UpdateResult } from 'typeorm';
import { SongsModel } from './songs.model';
import { InjectRepository } from '@nestjs/typeorm';
import { ArtistsModel } from 'src/artists/artists.model';
import { string } from 'joi';

@Injectable()
export class SongsService {
  constructor(
    @Inject('MYSQL_CONNECTION') mySqlConnection: MySqlConnectionType,
    @InjectRepository(SongsModel)
    private songsRepository: Repository<SongsModel>,
    @InjectRepository(ArtistsModel)
    private artistRepository: Repository<ArtistsModel>,
  ) {
    console.log(mySqlConnection); // this dummy example to show case useValue providers
  }

  async createSong(createSongPayload: CreateSongPayload): Promise<SongsModel> {
    const newSong = new SongsModel();
    newSong.title = createSongPayload.title;
    newSong.releasedDate = createSongPayload.releasedDate;
    newSong.duration = createSongPayload.duration;
    newSong.lyrics = createSongPayload.lyrics;

    const artists = await this.artistRepository.findByIds(
      createSongPayload.artists,
    );
    console.log('artists', artists);
    newSong.artists = artists;

    return this.songsRepository.save(newSong);
  }

  async findAll(page: number = 1, limit: number = 10): Promise<SongsModel[]> {
    const [songs] = await this.songsRepository.findAndCount({
      take: limit,
      skip: (page - 1) * limit,
    });
    return songs;
  }

  findOne(id: number): Promise<SongsModel> {
    return this.songsRepository.findOne({
      where: {
        id,
      },
    });
  }

  async updateOne(
    id: number,
    updateSongPayload: CreateSongPayload,
  ): Promise<UpdateResult> {
    const artistsFound = await this.artistRepository.findByIds(
      updateSongPayload.artists,
    );
    console.log(artistsFound);
    const artistsToUpdate = [];
    const { artists, ...rest } = updateSongPayload;
    if (artistsFound) {
      artistsFound.map((artists) => {
        artistsToUpdate.push({
          id: artists.id,
        });
      });
    }
    return this.songsRepository.update(id, {
      ...rest,
      artists: artistsToUpdate,
    });
  }

  deleteOne(id: number): Promise<DeleteResult> {
    return this.songsRepository.delete({
      id,
    });
  }
}
