import { AbstractRepository } from '@app/common/database/abstract.repository';
import { Logger } from '@nestjs/common';
import { ReservationsDocument } from './model/reservations.model';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

export class ReservationsRepository extends AbstractRepository<ReservationsDocument> {
  protected readonly logger = new Logger(ReservationsRepository.name);

  constructor(
    @InjectModel(ReservationsDocument.name)
    reservationModel: Model<ReservationsDocument>,
  ) {
    super(reservationModel);
  }
}
