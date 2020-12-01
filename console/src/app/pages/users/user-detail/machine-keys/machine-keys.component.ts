import { SelectionModel } from '@angular/cdk/collections';
import { Component, EventEmitter, Input, OnInit, Output, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { TranslateService } from '@ngx-translate/core';
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { Moment } from 'moment';
import { BehaviorSubject, Observable } from 'rxjs';
import { MachineKeySearchResponse, MachineKeyType, MachineKeyView } from 'src/app/proto/generated/management_pb';
import { ManagementService } from 'src/app/services/mgmt.service';
import { ToastService } from 'src/app/services/toast.service';

import { AddKeyDialogComponent } from './add-key-dialog/add-key-dialog.component';
import { ShowKeyDialogComponent } from './show-key-dialog/show-key-dialog.component';

@Component({
    selector: 'app-machine-keys',
    templateUrl: './machine-keys.component.html',
    styleUrls: ['./machine-keys.component.scss'],
})
export class MachineKeysComponent implements OnInit {
    @Input() userId!: string;
    @ViewChild(MatPaginator) public paginator!: MatPaginator;
    public dataSource: MatTableDataSource<MachineKeyView.AsObject> = new MatTableDataSource<MachineKeyView.AsObject>();
    public selection: SelectionModel<MachineKeyView.AsObject> = new SelectionModel<MachineKeyView.AsObject>(true, []);
    public keyResult!: MachineKeySearchResponse.AsObject;
    private loadingSubject: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    public loading$: Observable<boolean> = this.loadingSubject.asObservable();
    @Input() public displayedColumns: string[] = ['select', 'id', 'type', 'creationDate', 'expirationDate'];

    @Output() public changedSelection: EventEmitter<Array<MachineKeyView.AsObject>> = new EventEmitter();

    constructor(public translate: TranslateService, private userService: ManagementService, private dialog: MatDialog,
        private toast: ToastService) {
        this.selection.changed.subscribe(() => {
            this.changedSelection.emit(this.selection.selected);
        });
    }

    public ngOnInit(): void {
        this.getData(10, 0);
    }


    public isAllSelected(): boolean {
        const numSelected = this.selection.selected.length;
        const numRows = this.dataSource.data.length;
        return numSelected === numRows;
    }

    public masterToggle(): void {
        this.isAllSelected() ?
            this.selection.clear() :
            this.dataSource.data.forEach(row => this.selection.select(row));
    }


    public changePage(event: PageEvent): void {
        this.getData(event.pageSize, event.pageIndex * event.pageSize);
    }

    public deleteSelectedKeys(): void {
        const mappedDeletions = this.selection.selected.map(value => {
            return this.userService.DeleteMachineKey(value.id, this.userId);
        });
        Promise.all(mappedDeletions).then(() => {
            this.selection.clear();
            this.toast.showInfo('USER.TOAST.SELECTEDKEYSDELETED', true);
            this.getData(10, 0);
        }).catch(error => {
            this.toast.showError(error);
        });
    }

    public openAddKey(): void {
        const dialogRef = this.dialog.open(AddKeyDialogComponent, {
            data: {},
            width: '400px',
        });

        dialogRef.afterClosed().subscribe(resp => {
            if (resp) {
                const type: MachineKeyType = resp.type;

                let date: Timestamp | undefined;

                if (resp.date as Moment) {
                    const ts = new Timestamp();
                    console.log(resp.date.toDate());
                    const milliseconds = resp.date.toDate().getTime();
                    const seconds = Math.abs(milliseconds / 1000);
                    const nanos = (milliseconds - seconds * 1000) * 1000 * 1000;
                    ts.setSeconds(seconds);
                    ts.setNanos(nanos);
                    date = ts;
                }

                if (type) {
                    return this.userService.AddMachineKey(this.userId, type, date).then((response) => {
                        if (response) {
                            setTimeout(() => {
                                this.refreshPage();
                            }, 1000);

                            this.dialog.open(ShowKeyDialogComponent, {
                                data: {
                                    key: response.toObject(),
                                },
                                width: '400px',
                            });
                        }
                    }).catch((error: any) => {
                        this.toast.showError(error);
                    });
                }
            }
        });
    }

    private async getData(limit: number, offset: number): Promise<void> {
        this.loadingSubject.next(true);

        this.userService.SearchMachineKeys(this.userId, limit, offset).then(resp => {
            this.keyResult = resp.toObject();
            this.dataSource.data = this.keyResult.resultList;
            this.loadingSubject.next(false);
        }).catch((error: any) => {
            this.toast.showError(error);
            this.loadingSubject.next(false);
        });
    }

    public refreshPage(): void {
        this.getData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize);
    }
}
