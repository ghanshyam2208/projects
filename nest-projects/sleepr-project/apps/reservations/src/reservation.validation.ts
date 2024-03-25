import { BadRequestException, PipeTransform } from '@nestjs/common';
import {
  CreateReservationDto,
  CreateReservationSchema,
} from './dto/CreateReservationDto';

export class ReservationValidationPipe implements PipeTransform {
  public transform(value: CreateReservationDto): CreateReservationDto {
    const result = CreateReservationSchema.validate(value);
    if (result.error) {
      const errorMessages = result.error.details.map((d) => d.message).join();
      throw new BadRequestException(errorMessages);
    }
    return value;
  }
}
