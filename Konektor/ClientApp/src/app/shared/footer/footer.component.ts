import { Component } from '@angular/core';

@Component({
  selector: 'footer',
  templateUrl: './footer.component.html'
})
export class FooterComponent {
  public currentCount = 0;

  public incrementfooter() {
    this.currentCount++;
  }
}
