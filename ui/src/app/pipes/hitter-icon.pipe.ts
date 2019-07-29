import { Pipe, PipeTransform } from '@angular/core';
import { Hitters } from  '../hitters';
import { HttpClientJsonpModule } from '@angular/common/http';

@Pipe({
  name: 'hitterIcon',
  pure: true
})
export class HitterIconPipe implements PipeTransform {

  transform(value: any, args?: any): any {
    return this.getHitterIcon(value);
  }

  getHitterIcon(hitter: number): string {
    // logic mostly taken from Base/Rules/CommonScripts/KillMessages.as
    switch(hitter) {
      case Hitters.nothing:
      case Hitters.suicide:
        return Hitters[Hitters.fall];
      case Hitters.fire:
      case Hitters.burn:
        return Hitters[Hitters.fire];
      case Hitters.bomb_arrow:
      case Hitters.explosion:
        return Hitters[Hitters.bomb];
      case Hitters.mine_special:
        return Hitters[Hitters.mine];
      case Hitters.cata_boulder:
        return Hitters[Hitters.boulder];
    }

    return Hitters[hitter];
  }

}
