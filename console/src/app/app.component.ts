import { BreakpointObserver } from '@angular/cdk/layout';
import { OverlayContainer } from '@angular/cdk/overlay';
import { DOCUMENT, ViewportScroller } from '@angular/common';
import { Component, ElementRef, HostBinding, Inject, OnDestroy, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatIconRegistry } from '@angular/material/icon';
import { MatDrawer } from '@angular/material/sidenav';
import { DomSanitizer } from '@angular/platform-browser';
import { Router, RouterOutlet } from '@angular/router';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { BehaviorSubject, from, Observable, of, Subscription } from 'rxjs';
import { catchError, debounceTime, finalize, map, take } from 'rxjs/operators';

import { accountCard, adminLineAnimation, navAnimations, routeAnimations, toolbarAnimation } from './animations';
import {
    MyProjectOrgSearchKey,
    MyProjectOrgSearchQuery,
    Org,
    SearchMethod,
    UserProfileView,
} from './proto/generated/auth_pb';
import { AuthenticationService } from './services/authentication.service';
import { GrpcAuthService } from './services/grpc-auth.service';
import { ManagementService } from './services/mgmt.service';
import { ThemeService } from './services/theme.service';
import { UpdateService } from './services/update.service';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    animations: [
        toolbarAnimation,
        ...navAnimations,
        accountCard,
        routeAnimations,
        adminLineAnimation,
    ],
})
export class AppComponent implements OnDestroy {
    @ViewChild('drawer') public drawer!: MatDrawer;
    @ViewChild('input', { static: false }) input!: ElementRef;
    public isHandset$: Observable<boolean> = this.breakpointObserver
        .observe('(max-width: 599px)')
        .pipe(map(result => {
            return result.matches;
        }));
    @HostBinding('class') public componentCssClass: string = 'dark-theme';

    public showAccount: boolean = false;
    public org!: Org.AsObject;
    public orgs$: Observable<Org.AsObject[]> = of([]);
    public profile!: UserProfileView.AsObject;
    public isDarkTheme: Observable<boolean> = of(true);

    public orgLoading$: BehaviorSubject<any> = new BehaviorSubject(false);

    public showProjectSection: boolean = false;

    public filterControl: FormControl = new FormControl('');
    private authSub: Subscription = new Subscription();
    private orgSub: Subscription = new Subscription();

    public hideAdminWarn: boolean = true;
    constructor(
        public viewPortScroller: ViewportScroller,
        @Inject('windowObject') public window: Window,
        public translate: TranslateService,
        public authenticationService: AuthenticationService,
        public authService: GrpcAuthService,
        private breakpointObserver: BreakpointObserver,
        public overlayContainer: OverlayContainer,
        private themeService: ThemeService,
        public mgmtService: ManagementService,
        public matIconRegistry: MatIconRegistry,
        public domSanitizer: DomSanitizer,
        private router: Router,
        update: UpdateService,
        @Inject(DOCUMENT) private document: Document,
    ) {
        console.log('%cWait!', 'text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black; color: #5469D4; font-size: 50px');
        console.log('%cInserting something here could give attackers access to your zitadel account.', 'color: red; font-size: 18px');
        console.log('%cIf you don\'t know exactly what you\'re doing, close the window and stay on the safe side', 'font-size: 16px');
        console.log('%cIf you know exactly what you are doing, you should work for us', 'font-size: 16px');
        this.setLanguage();

        this.matIconRegistry.addSvgIcon(
            'mdi_account_check_outline',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/account-check-outline.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_account_cancel',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/account-cancel-outline.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_light_on',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/lightbulb-on-outline.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_content_copy',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/content-copy.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_light_off',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/lightbulb-off-outline.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_radar',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/radar.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_lock_question',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/lock-question.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_textbox_password',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/textbox-password.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_lock_reset',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/lock-reset.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_broom',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/broom.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_pin_outline',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/pin-outline.svg'),
        );

        this.matIconRegistry.addSvgIcon(
            'mdi_pin',
            this.domSanitizer.bypassSecurityTrustResourceUrl('assets/mdi/pin.svg'),
        );
        this.getProjectCount();

        this.orgSub = this.authService.activeOrgChanged.subscribe(org => {
            this.org = org;
            this.getProjectCount();
        });

        this.authSub = this.authenticationService.authenticationChanged.subscribe((authenticated) => {
            if (authenticated) {
                this.authService.GetActiveOrg().then(org => {
                    this.org = org;
                });
            }
        });

        const theme = localStorage.getItem('theme');
        if (theme) {
            this.overlayContainer.getContainerElement().classList.add(theme);
            this.componentCssClass = theme;
        }

        this.isDarkTheme = this.themeService.isDarkTheme;
        this.isDarkTheme.subscribe(thema => this.onSetTheme(thema ? 'dark-theme' : 'light-theme'));

        this.translate.onLangChange.subscribe((language: LangChangeEvent) => {
            this.document.documentElement.lang = language.lang;
        });

        this.filterControl.valueChanges.pipe(debounceTime(300)).subscribe(value => {
            this.loadOrgs(
                value.trim().toLowerCase(),
            );
        });

        this.hideAdminWarn = localStorage.getItem('hideAdministratorWarning') === 'true' ? true : false;
    }

    public ngOnDestroy(): void {
        this.authSub.unsubscribe();
        this.orgSub.unsubscribe();
    }

    public toggleAdminHide(): void {
        this.hideAdminWarn = !this.hideAdminWarn;
        localStorage.setItem('hideAdministratorWarning', this.hideAdminWarn.toString());
    }

    public loadOrgs(filter?: string): void {
        let query;
        if (filter) {
            query = new MyProjectOrgSearchQuery();
            query.setMethod(SearchMethod.SEARCHMETHOD_CONTAINS_IGNORE_CASE);
            query.setKey(MyProjectOrgSearchKey.MYPROJECTORGSEARCHKEY_ORG_NAME);
            query.setValue(filter);
        }

        this.orgLoading$.next(true);
        this.orgs$ = from(this.authService.SearchMyProjectOrgs(10, 0, query ? [query] : undefined)).pipe(
            map(resp => {
                return resp.toObject().resultList;
            }),
            catchError(() => of([])),
            finalize(() => {
                this.orgLoading$.next(false);
                this.focusFilter();
            }),
        );
    }

    public prepareRoute(outlet: RouterOutlet): boolean {
        return outlet && outlet.activatedRouteData && outlet.activatedRouteData.animation;
    }

    public closeAccountCard(): void {
        if (this.showAccount) {
            this.showAccount = false;
        }
    }

    public onSetTheme(theme: string): void {
        localStorage.setItem('theme', theme);
        this.overlayContainer.getContainerElement().classList.add(theme);
        this.componentCssClass = theme;
    }

    private setLanguage(): void {
        this.translate.addLangs(['en', 'de']);
        this.translate.setDefaultLang('en');

        this.authService.user.subscribe(userprofile => {
            this.profile = userprofile;
            const cropped = navigator.language.split('-')[0] ?? 'en';
            const fallbackLang = cropped.match(/en|de/) ? cropped : 'en';
            const lang = userprofile.preferredLanguage.match(/en|de/) ? userprofile.preferredLanguage : fallbackLang;
            this.translate.use(lang);
            this.document.documentElement.lang = lang;
        });
    }

    public setActiveOrg(org: Org.AsObject): void {
        this.org = org;
        this.authService.setActiveOrg(org);
        this.authService.zitadelPermissionsChanged.pipe(take(1)).subscribe(() => {
            this.router.navigate(['/']);
        });
    }

    private getProjectCount(): void {
        this.authService.isAllowed(['project.read$']).subscribe((allowed) => {
            if (allowed) {
                this.mgmtService.SearchProjects(0, 0);

                this.mgmtService.SearchGrantedProjects(0, 0);
            }
        });
    }

    focusFilter(): void {
        setTimeout(() => {
            this.input.nativeElement.focus();
        }, 0);
    }
}

