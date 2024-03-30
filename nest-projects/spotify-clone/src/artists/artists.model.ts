import { SongsModel } from 'src/songs/songs.model';
import { UsersModel } from 'src/users/users.model';
import {
  Entity,
  JoinColumn,
  ManyToMany,
  OneToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('Artists')
export class ArtistsModel {
  @PrimaryGeneratedColumn()
  id: number;

  @OneToOne(() => UsersModel)
  @JoinColumn()
  user: UsersModel;

  @ManyToMany(() => SongsModel, (song) => song.artists)
  songs: SongsModel[];
}
