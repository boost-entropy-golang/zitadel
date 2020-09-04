import { animate, animateChild, keyframes, query, stagger, style, transition, trigger } from '@angular/animations';
import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { CreationType, MemberCreateDialogComponent } from 'src/app/modules/add-member-dialog/member-create-dialog.component';
import { MemberType, UserMembershipSearchResponse, UserView } from 'src/app/proto/generated/management_pb';
import { AdminService } from 'src/app/services/admin.service';
import { ManagementService } from 'src/app/services/mgmt.service';
import { ToastService } from 'src/app/services/toast.service';

@Component({
    selector: 'app-memberships',
    templateUrl: './memberships.component.html',
    styleUrls: ['./memberships.component.scss'],
    animations: [
        trigger('cardAnimation', [
            transition('* => *', [
                query('@animate', stagger('40ms', animateChild()), { optional: true }),
            ]),
        ]),
        trigger('animate', [
            transition(':enter', [
                animate('.2s ease-in', keyframes([
                    style({ opacity: 0, offset: 0 }),
                    style({ opacity: .5, transform: 'scale(1.05)', offset: 0.3 }),
                    style({ opacity: 1, transform: 'scale(1)', offset: 1 }),
                ])),
            ]),
        ]),
    ],
})
export class MembershipsComponent implements OnInit {
    public loading: boolean = false;
    public memberships!: UserMembershipSearchResponse.AsObject;

    @Input() public user!: UserView.AsObject;
    @Input() public disabled: boolean = false;

    public MemberType: any = MemberType;

    constructor(
        private mgmtService: ManagementService,
        private adminService: AdminService,
        private dialog: MatDialog,
        private toast: ToastService,
        private router: Router,
    ) { }

    ngOnInit(): void {
        this.loadManager(this.user.id);
        console.log(this.disabled);
    }

    public async loadManager(userId: string): Promise<void> {
        this.mgmtService.SearchUserMemberships(userId, 100, 0, []).then(response => {
            this.memberships = response.toObject();
            this.loading = false;
        });
    }

    public navigateToObject(): void {
        if (!this.disabled) {
            this.router.navigate(['/users', this.user.id, 'memberships']);
        }
    }

    public addMember(): void {
        const dialogRef = this.dialog.open(MemberCreateDialogComponent, {
            width: '400px',
            data: {
                user: this.user,
            },
        });

        dialogRef.afterClosed().subscribe(resp => {
            if (resp && resp.creationType !== undefined) {
                switch (resp.creationType) {
                    case CreationType.IAM:
                        this.createIamMember(resp);
                        break;
                    case CreationType.ORG:
                        this.createOrgMember(resp);
                        break;
                    case CreationType.PROJECT_OWNED:
                        this.createOwnedProjectMember(resp);
                        break;
                    case CreationType.PROJECT_GRANTED:
                        this.createGrantedProjectMember(resp);
                        break;
                }
            }
        });
    }

    public createIamMember(response: any): void {
        const users: UserView.AsObject[] = response.users;
        const roles: string[] = response.roles;

        if (users && users.length && roles && roles.length) {
            Promise.all(users.map(user => {
                return this.adminService.AddIamMember(user.id, roles);
            })).then(() => {
                this.toast.showInfo('IAM.TOAST.MEMBERADDED', true);
            }).catch(error => {
                this.toast.showError(error);
            });
        }
    }

    private createOrgMember(response: any): void {
        const users: UserView.AsObject[] = response.users;
        const roles: string[] = response.roles;

        if (users && users.length && roles && roles.length) {
            Promise.all(users.map(user => {
                return this.mgmtService.AddMyOrgMember(user.id, roles);
            })).then(() => {
                this.toast.showInfo('ORG.TOAST.MEMBERADDED', true);
            }).catch(error => {
                this.toast.showError(error);
            });
        }
    }

    private createGrantedProjectMember(response: any): void {
        const users: UserView.AsObject[] = response.users;
        const roles: string[] = response.roles;

        if (users && users.length && roles && roles.length) {
            users.forEach(user => {
                return this.mgmtService.AddProjectGrantMember(
                    response.projectId,
                    response.grantId,
                    user.id,
                    roles,
                ).then(() => {
                    this.toast.showInfo('PROJECT.TOAST.MEMBERADDED', true);
                }).catch(error => {
                    this.toast.showError(error);
                });
            });
        }
    }

    private createOwnedProjectMember(response: any): void {
        const users: UserView.AsObject[] = response.users;
        const roles: string[] = response.roles;

        if (users && users.length && roles && roles.length) {
            users.forEach(user => {
                return this.mgmtService.AddProjectMember(response.projectId, user.id, roles)
                    .then(() => {
                        this.toast.showInfo('PROJECT.TOAST.MEMBERADDED', true);
                    }).catch(error => {
                        this.toast.showError(error);
                    });
            });
        }
    }

    getColor(type: MemberType): string {
        const gen = type.toString();
        const colors = [
            'rgb(201, 115, 88)',
            'rgb(226, 176, 50)',
            'rgb(112, 89, 152)',
        ];

        let hash = 0;
        if (gen.length === 0) {
            return colors[hash];
        }
        for (let i = 0; i < gen.length; i++) {
            // tslint:disable-next-line: no-bitwise
            hash = gen.charCodeAt(i) + ((hash << 5) - hash);
            // tslint:disable-next-line: no-bitwise
            hash = hash & hash;
        }
        hash = ((hash % colors.length) + colors.length) % colors.length;
        return colors[hash];
    }
}
