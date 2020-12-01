import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatTooltipModule } from '@angular/material/tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { HasRoleModule } from 'src/app/directives/has-role/has-role.module';
import { CardModule } from 'src/app/modules/card/card.module';
import { DetailLayoutModule } from 'src/app/modules/detail-layout/detail-layout.module';
import { IdpTableModule } from 'src/app/modules/idp-table/idp-table.module';
import { InputModule } from 'src/app/modules/input/input.module';
import { HasRolePipeModule } from 'src/app/pipes/has-role-pipe/has-role-pipe.module';

import { AddIdpDialogModule } from './add-idp-dialog/add-idp-dialog.module';
import { LoginPolicyRoutingModule } from './login-policy-routing.module';
import { LoginPolicyComponent } from './login-policy.component';

@NgModule({
    declarations: [LoginPolicyComponent],
    imports: [
        LoginPolicyRoutingModule,
        CommonModule,
        FormsModule,
        CardModule,
        InputModule,
        MatButtonModule,
        MatSlideToggleModule,
        MatIconModule,
        HasRoleModule,
        HasRolePipeModule,
        MatTooltipModule,
        TranslateModule,
        DetailLayoutModule,
        AddIdpDialogModule,
        IdpTableModule,
        MatProgressSpinnerModule,
    ],
})
export class LoginPolicyModule { }
