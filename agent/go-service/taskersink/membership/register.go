package membership

import maa "github.com/MaaXYZ/maa-framework-go/v4"

// Register registers the membership check custom action for pipeline-level gating.
func Register() {
	maa.AgentServerRegisterCustomAction("MembershipCheck", &MembershipCheckAction{})
}
