import { Component, HostBinding } from '@angular/core';

@Component({
  selector: 'chef-table-header-cell',
  templateUrl: './table-header-cell.component.html',
  styleUrls: ['./table-header-cell.component.scss']
})
export class TableHeaderCellComponent {
  @HostBinding('attr.role') role = 'columnheader';
}
