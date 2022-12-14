import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ListInputComponent } from './list-input.component';

describe('ListInputComponent', () => {
  let component: ListInputComponent;
  let fixture: ComponentFixture<ListInputComponent>;

  beforeEach(
    waitForAsync(() => {
      TestBed.configureTestingModule({
        declarations: [ListInputComponent],
      }).compileComponents();
    }),
  );

  beforeEach(() => {
    fixture = TestBed.createComponent(ListInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
