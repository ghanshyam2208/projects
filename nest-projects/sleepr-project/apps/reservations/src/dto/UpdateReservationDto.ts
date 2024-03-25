import { PartialType } from '@nestjs/mapped-types';
import { CreateReservationDto } from './CreateReservationDto';

export class UpdateReservationDto extends PartialType(CreateReservationDto) {}
