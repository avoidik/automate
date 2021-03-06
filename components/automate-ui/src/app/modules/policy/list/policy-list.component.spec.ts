import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { Store, StoreModule } from '@ngrx/store';
import { NgrxStateAtom, ngrxReducers, defaultInitialState, runtimeChecks } from 'app/ngrx.reducers';
import { of as observableOf } from 'rxjs';
import { MockComponent } from 'ng2-mock-component';

import { ChefPipesModule } from 'app/pipes/chef-pipes.module';
import { customMatchers } from 'app/testing/custom-matchers';
import { PolicyListComponent } from './policy-list.component';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';

describe('PolicyListComponent', () => {
  let store: Store<NgrxStateAtom>;
  let component: PolicyListComponent;
  let fixture: ComponentFixture<PolicyListComponent>;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        MockComponent({ selector: 'app-authorized', inputs: ['allOf'] }),
        MockComponent({
          selector: 'app-delete-object-modal',
          inputs: ['visible', 'objectNoun', 'objectName', 'moreDetails'],
          outputs: ['close', 'deleteClicked']
        }),
        MockComponent({ selector: 'chef-control-menu' }),
        MockComponent({ selector: 'chef-heading' }),
        MockComponent({ selector: 'chef-option' }),
        MockComponent({ selector: 'chef-page-header' }),
        MockComponent({ selector: 'chef-subheading' }),
        MockComponent({ selector: 'chef-loading-spinner' }),
        MockComponent({ selector: 'chef-table-new' }),
        MockComponent({ selector: 'chef-table-body' }),
        MockComponent({ selector: 'chef-table-cell' }),
        MockComponent({ selector: 'chef-table-header-cell' }),
        MockComponent({ selector: 'chef-table-header' }),
        MockComponent({ selector: 'chef-table-row' }),
        MockComponent({ selector: 'a', inputs: ['routerLink'] }),
        PolicyListComponent
      ],
      providers: [
        FeatureFlagsService
      ],
      imports: [
        ChefPipesModule,
        StoreModule.forRoot(ngrxReducers, { ...defaultInitialState, runtimeChecks })
      ]
    }).compileComponents();
      store = TestBed.get(Store);
      spyOn(store, 'dispatch').and.callThrough();
  }));

  beforeEach(() => {
    jasmine.addMatchers(customMatchers);
    fixture = TestBed.createComponent(PolicyListComponent);
    component = fixture.componentInstance;
    element = fixture.debugElement.nativeElement;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('contains key elements', () => {
    expect(element).toContainPath('chef-page-header');
  });

  it('displays policy data for v2', () => {
    component.isIAMv2$ = observableOf(true);
    fixture.detectChanges();
    expect(element).toContainPath('app-authorized');
  });

  it('does not display policy data for v1', () => {
    component.isIAMv2$ = observableOf(false);
    fixture.detectChanges();
    expect(element).not.toContainPath('app-authorized');
  });
});
