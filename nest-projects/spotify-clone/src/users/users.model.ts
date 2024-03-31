import { PlaylistsModel } from 'src/playlists/playlists.model';
import { Column, Entity, OneToMany, PrimaryGeneratedColumn } from 'typeorm';

@Entity('Users')
export class UsersModel {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  firstName: string;

  @Column()
  lastName: string;

  @Column()
  email: string;

  @Column()
  password: string;

  @OneToMany(() => PlaylistsModel, (playlist) => playlist.user)
  playlists: PlaylistsModel[];
}
