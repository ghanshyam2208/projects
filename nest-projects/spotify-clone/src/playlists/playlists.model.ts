import { SongsModel } from 'src/songs/songs.model';
import { UsersModel } from 'src/users/users.model';
import {
  Column,
  Entity,
  ManyToOne,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('playlists')
export class PlaylistsModel {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @OneToMany(() => SongsModel, (song) => song.playlist)
  songs: SongsModel[];

  @ManyToOne(() => UsersModel, (user) => user.playlists)
  user: UsersModel;
}
