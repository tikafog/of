package of

const (
	// NameNotSet is a Name of type when you don't specify a name
	NameNotSet Name = "notset"
	// NameAccount is a Name of type Account.
	NameAccount Name = "account"
	// NameAdmin is a Name of type Admin.
	NameAdmin Name = "admin"
	// NameAdminClient is a Name of type Admin Client.
	NameAdminClient Name = "admin_client"
	// NameCenter is a Name of type Center.
	NameCenter Name = "center"
	// NameNode is a Name of type Node.
	NameNode Name = "node"
	// NameBootNode is a Name of type BootNode.
	NameBootNode Name = "bootnode"
	// NameUpgrade is a Name of type Upgrade.
	NameUpgrade Name = "upgrade"
	// NameMedia is a Name of type Media.
	NameMedia Name = "media"
	// NameKernel is a Name of type Kernel.
	NameKernel Name = "kernel"
	// NameGateway is a Name of type Gateway.
	NameGateway Name = "gateway"
)

var SupportModuleNames = []Name{
	NameAccount,
	NameAdmin,
	NameAdminClient,
	NameCenter,
	NameNode,
	NameBootNode,
	NameUpgrade,
	NameMedia,
	NameKernel,
	NameGateway,
}
