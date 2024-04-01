import { Injectable } from '@nestjs/common';
import { CreateSongPayload } from './songs.validations';
import { SongsRepository } from './songs.repo';

@Injectable()
export class SongsService {
  constructor(private readonly songsRepository: SongsRepository) {}
  createSong(createSongPayload: CreateSongPayload) {
    return this.songsRepository.createSong(createSongPayload);
  }
}
