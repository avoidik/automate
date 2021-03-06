import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { ChefComponentsModule } from 'app/components/chef-components.module';

import { UserDetailsComponent } from './user-details/user-details.component';
import { UserManagementComponent } from './user-management/user-management.component';
import { UserFormComponent } from './user-management/user-form/user-form.component';
import { UserTableComponent } from './user-table/user-table.component';
import { UserDetailsNonAdminResolve } from './user-details/user-details.resolver';
import { UserProfileSidebarComponent } from './user-profile-sidebar/user-profile-sidebar.component';

import { UserRoutingModule } from './user-routing.module';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ChefComponentsModule,
    UserRoutingModule
  ],
  exports: [
    UserDetailsComponent,
    UserManagementComponent,
    UserTableComponent
  ],
  declarations: [
    UserDetailsComponent,
    UserManagementComponent,
    UserFormComponent,
    UserTableComponent,
    UserProfileSidebarComponent
  ],
  providers: [
    UserDetailsNonAdminResolve
  ],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
})
export class UserModule { }
