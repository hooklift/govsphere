//Copyright (c) 2014, Cloudescape. All rights reserved.

/*
The VMware vSphere management object model is a complex system of
data structures designed to provision, manage, monitor, and control
the life-cycle of all components that can possibly comprise
virtual infrastructure. The VMware vSphere management architecture
is patterned after Java’s JMX (Java Management Extensions)
infrastructure in which objects are used to instrument other objects,
on a remote server. The data structures defined for the object
model include both managed object types, as described on this page,
and data object types.

Managed Object Type

A “managed object type” is a core data structure of the server-side
object model. Instances of various managed object types are referred
to generically as “managed objects,” of which there are two
broad categories:

• Managed objects that extend the ManagedEntity managed object type,
and thus, are components that comprise the inventory of virtual components.
For example, instances of host systems (HostSystem),
virtual machines (VirtualMachine, and datastores (Datastore) are
inventoried objects, and are referred to generically as “managed entities.”

• Managed objects that provide services for the entire system.
Managed objects in this category enable managing performance
(PerformanceManager), managing licenses for VMware products
(LicenseManager), and managing virtual storage (VirtualDiskManager).
These managed objects are the service interfaces for the
virtual infrastructure management components.

Managed objects can contain both properties and operations.
An “operation” is Web-services terminology for what might be
called a “method” in other programming languages, such as Java.
In fact, the word “method” is used in the API Reference rather than
operation, but you may see the two words used, interchangeably.

Regardless of these subtle language differences, working with the
server from a client involves a few common steps, starting with
connecting to the server, authenticating user-account credentials,
and obtaining a session.

After connecting to the server system, the client application
must then obtain a reference to the ServiceInstance managed object.
*/
package mo
