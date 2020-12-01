import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatChipsModule } from '@angular/material/chips';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatRadioModule } from '@angular/material/radio';
import { MatSelectModule } from '@angular/material/select';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatStepperModule } from '@angular/material/stepper';
import { MatTooltipModule } from '@angular/material/tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { CopyToClipboardModule } from 'src/app/directives/copy-to-clipboard/copy-to-clipboard.module';
import { HasRoleModule } from 'src/app/directives/has-role/has-role.module';
import { CardModule } from 'src/app/modules/card/card.module';
import { ChangesModule } from 'src/app/modules/changes/changes.module';
import { InfoSectionModule } from 'src/app/modules/info-section/info-section.module';
import { InputModule } from 'src/app/modules/input/input.module';
import { MetaLayoutModule } from 'src/app/modules/meta-layout/meta-layout.module';

import { AppCreateComponent } from './app-create/app-create.component';
import { AppDetailComponent } from './app-detail/app-detail.component';
import { AppSecretDialogComponent } from './app-secret-dialog/app-secret-dialog.component';
import { AppsRoutingModule } from './apps-routing.module';

@NgModule({
    declarations: [
        AppCreateComponent,
        AppDetailComponent,
        AppSecretDialogComponent,
    ],
    imports: [
        CommonModule,
        AppsRoutingModule,
        FormsModule,
        TranslateModule,
        ReactiveFormsModule,
        HasRoleModule,
        MatMenuModule,
        MatChipsModule,
        MatIconModule,
        MatSelectModule,
        MatButtonToggleModule,
        MatButtonModule,
        MatProgressSpinnerModule,
        MatProgressBarModule,
        MatDialogModule,
        MatCheckboxModule,
        CardModule,
        MatTooltipModule,
        TranslateModule,
        MatStepperModule,
        MatRadioModule,
        CopyToClipboardModule,
        MatSlideToggleModule,
        InputModule,
        MetaLayoutModule,
        ChangesModule,
        InfoSectionModule,
    ],
    exports: [TranslateModule],
})
export class AppsModule { }
