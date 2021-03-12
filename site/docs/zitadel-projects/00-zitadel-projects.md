---
title: ZITADEL Projects
tags: beginner, administrator, product owner
readingtime: 10min
---


| | |
| --- | --- |
| Description | Learn the basics about applications, roles and authorizations, and how projects allow you to group these together. |
| Learning Outcomes | In this module you will: <ul><li>Learn about projects and granted projects</li><li>Create a new project</li><li>Creating simple roles and authorizations</li><li>Create an organization grant for your project</li></ul> |
| Prerequisites | <ul><li>ZITADEL organizations</li><li>Role Based Access Management (RBAC)</li></ul> |

### What is a project?

The idea of projects is to have a vessel for all components who are closely related to each other. Multiple projects can exist within an organization.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/zitadel_organization_grant.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/zitadel_organization_grant.png" itemprop="thumbnail" alt="Zitadel Organization grant" />
        </a>
        <figcaption itemprop="caption description">Organization Grant</figcaption>
    </figure>
</div>

All applications within a project share the same roles, grants, and authorizations:

* **Applications** is your software that initiates the authorization flow. This could be a web app and a mobile app that share the same roles.
* **Roles** are a means of managing user access rights for a project. 
* **Authorizations** define which users have which roles. Authorizations are also called “user grants”.
* **Granted Organizations** can manage selected roles for your project on their own.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_overview.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_overview.png" itemprop="thumbnail" alt="console_projects_overview.png" />
        </a>
        <figcaption itemprop="caption description">Organization Grant</figcaption>
    </figure>
</div>

The goal of this module is to give you an overview, but not dive too deep into details around managing access rights and delegating management of roles to third parties. So let’s create a straightforward example project first.

### Exercise - Create a simple project

Visit https://console.zitadel.ch/projects or select “Projects” within your organization, then click the button to create a new project. 

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_empty.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_empty.png" itemprop="thumbnail" alt="console_projects_empty.png" />
        </a>
        <figcaption itemprop="caption description">Empty Project</figcaption>
    </figure>
</div>

Enter the name “ My first project” and continue.

Let’s make this more interesting and add some basic roles and authorizations to your project and then confirm the scope of the roles and authorizations.

Jump to the section ROLES and create two new roles with the following values

* Key: reader
* Display Name: Reader
* Group: user

and

* Key: editor
* Display Name: Editor
* Group: user

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_add_new_roles.gif" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_add_new_roles.gif" itemprop="thumbnail" alt="console_projects_add_new_roles" />
        </a>
        <figcaption itemprop="caption description">Add New Roles</figcaption>
    </figure>
</div>

Now, you can add roles to your own user, or you can create a new user. To create a new user, go to Users and click “New”. Enter the required contact details and save by clicking “Create”.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_users_create_new_user.gif" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_users_create_new_user.gif" itemprop="thumbnail" alt="console_users_create_new_user" />
        </a>
        <figcaption itemprop="caption description">Create new user</figcaption>
    </figure>
</div>

To grant users certain roles, you need to create authorizations. Go back to the project, and jump to the section AUTHORIZATIONS.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_create_authorization.gif" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_create_authorization.gif" itemprop="thumbnail" alt="console_projects_create_authorization" />
        </a>
        <figcaption itemprop="caption description">Create Authorization</figcaption>
    </figure>
</div>

You can verify the role grant on the user. Select Users from the navigation menu and click on the user Coyote. Scroll down to the section AUTHORIZATION, there you should be able to verify that the user has the role ‘reader’ for your project ‘My first project’.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_authorization_created.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_authorization_created.png" itemprop="thumbnail" alt="console_projects_authorization_created" />
        </a>
        <figcaption itemprop="caption description">Verify your authorization</figcaption>
    </figure>
</div>


Now create another project (eg. “My second project”) and verify that there are no roles or authorizations on your second project.

### What is a granted project?

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/zitadel_organization_grant.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/zitadel_organization_grant.png" itemprop="thumbnail" alt="Zitadel Organization grant" />
        </a>
        <figcaption itemprop="caption description">Organization Grant</figcaption>
    </figure>
</div>

With ZITADEL you can grant selected roles within your project to an organization. The receiving organization can then create authorizations for their users on their own (self-service). 

An example could be a service provider that offers a SaaS solution that has different permissions for employees working in Sales and Accounting. As soon as a new client purchases the service, the provider could grant the roles ‘sales’ and ‘accounting’ to that organization, allowing them to self-manage how they want to allocate the roles to their users.

The process of assigning certain roles by default or according to rules can be further automated by utilizing Service Users in the service provider’s business process.

Obviously, your organization can grant projects and receive projects. When you are granting, then the receiving organization will be displayed in the section GRANTED ORGANIZATIONS of your project.

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_projects_granted.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_projects_granted.png" itemprop="thumbnail" alt="console_projects_granted" />
        </a>
        <figcaption itemprop="caption description">Project granted to organization</figcaption>
    </figure>
</div>

A granted project, on the other hand, belongs to a third party, granting you some rights to manage certain roles of their project. ZITADEL Console shows granted projects in a dedicated navigation menu, to clearly separate from your organization’s projects. 

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/console_granted_projects_overview.png" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/console_granted_projects_overview.png" itemprop="thumbnail" alt="console_granted_projects_overview" />
        </a>
        <figcaption itemprop="caption description">Granted Projects Overview</figcaption>
    </figure>
</div>

Please note that you can also grant selected roles of a project to an individual user, instead of an organization. We will discuss this in more detail in a later section.

### Exercise - Grant a project

1. Visit the project that you have created before, then in the section GRANTED ORGANIZATIONS click New. 
2. Enter the domain ‘acme.zitadel.ch’, search the organization and continue to the next step.
3. Select some roles you would like to grant to the organization ACME and confirm.
4. You should now see ACME in the section GRANTED ORGANIZATIONS

<div class="zitadel-gallery" itemscope itemtype="http://schema.org/ImageGallery">
    <figure itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">
        <a href="img/projects_create_org_grant_caos2acme.gif" itemprop="contentUrl" data-size="1920x1080">
            <img src="img/projects_create_org_grant_caos2acme.gif" itemprop="thumbnail" alt="projects_create_org_grant_caos2acme" />
        </a>
        <figcaption itemprop="caption description">Grant a project</figcaption>
    </figure>
</div>


### Knowledge Check (2)

* You can setup multiple projects within an organization to manage scope
    - [ ] yes
    - [ ] no
* Authorizations are define more detailed access rights within your application
    - [ ] yes
    - [ ] no
* Your projects as well as projects granted to your organization are visible within the Tab Projects of your organization
    - [ ] yes
    - [ ] no

<details>
    <summary>
        Solutions
    </summary>

* You can setup multiple projects within an organization to manage scope
    - [x] yes
    - [ ] no
* Authorizations are define more detailed access rights within your application
    - [ ] yes
    - [x] no (Authorizations link users to certain roles)
* Your projects as well as projects granted to your organization are visible within the Tab Projects of your organization
    - [ ] yes
    - [x] no (Projects and Granted Projects are shown on different tabs)
    
</details>

### Summary (2)

* Manage scope of roles, authorizations and applications with projects
* Create and assign roles to users of your organization within your project
* Use project grants to enable other organizations to self-manage access rights (roles) to your applications

Where to go from here: 
* Manage roles for your project
* Grant roles to other organizations or users
* Service Users
* Manage IAM Roles
* Setup a SaaS Application with granted projects (Learning Path)