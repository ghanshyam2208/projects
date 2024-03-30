import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@Entity('playlists')
export class PlaylistsModel {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;
}
