package shoot

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func suppressMissingOptionalConfigurationBlock(k, old, new string, d *schema.ResourceData) bool {
	return old == "1" && new == "0"
}

func suppressEmptyNewValue(k, old, new string, d *schema.ResourceData) bool {
	return len(new) == 0
}

func suppressStatusLabel(k, old, new string, d *schema.ResourceData) bool {
	return strings.HasSuffix(k, "shoot.garden.sapcloud.io/status") || strings.HasSuffix(k, "labels.%")
}

func suppressCreatedByAnnotation(k, old, new string, d *schema.ResourceData) bool {
	return strings.HasSuffix(k, "garden.sapcloud.io/createdBy") || strings.HasSuffix(k, "annotations.%")
}