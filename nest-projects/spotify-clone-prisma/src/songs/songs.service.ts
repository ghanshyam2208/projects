import { Injectable } from '@nestjs/common';
import { SongsRepository } from './songs.repository';
import { CreateSongPayload } from './songs.validations';

@Injectable()
export class SongsService {
  constructor(private readonly songsRepository: SongsRepository) {}

  createSong(createSongPayload: CreateSongPayload) {
    return this.songsRepository.createSong(createSongPayload);
  }
}
