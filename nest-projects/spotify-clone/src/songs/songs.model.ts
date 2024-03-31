import { ArtistsModel } from 'src/artists/artists.model';
import { PlaylistsModel } from 'src/playlists/playlists.model';
import {
  Column,
  Entity,
  JoinTable,
  ManyToMany,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('Songs')
export class SongsModel {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  title: string;

  // @Column('varchar', { array: true })
  // artists: string[];

  @Column('date')
  releasedDate: Date;

  @Column('time')
  duration: Date;

  @Column('text')
  lyrics: string;

  @ManyToMany(() => ArtistsModel, (artist) => artist.songs, { cascade: true })
  @JoinTable({ name: 'songs_artists' })
  artists: ArtistsModel[];

  @ManyToOne(() => PlaylistsModel, (playlist) => playlist.songs)
  playlist: PlaylistsModel;
}
