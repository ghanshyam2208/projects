import { UsersModel } from 'src/users/users.model';
import { Entity, JoinColumn, OneToOne, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class ArtistsModel {
  @PrimaryGeneratedColumn()
  id: number;

  @OneToOne(() => UsersModel)
  @JoinColumn()
  user: UsersModel;
}
