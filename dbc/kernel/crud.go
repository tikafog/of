// Code generated by ent, DO NOT EDIT.

package kernel

func (ic *InstructCreate) SetInstruct(input *Instruct) *InstructCreate {
	ic.SetCurrentUnix(input.CurrentUnix)
	ic.SetUpdatedUnix(input.UpdatedUnix)
	return ic
}

func (ic *InstructCreate) SetInstructWithOptional(input *Instruct) *InstructCreate {
	ic.SetCurrentUnix(input.CurrentUnix)
	ic.SetUpdatedUnix(input.UpdatedUnix)
	return ic
}

func (iuo *InstructUpdateOne) SetInstruct(input *Instruct) *InstructUpdateOne {
	iuo.SetCurrentUnix(input.CurrentUnix)
	iuo.SetUpdatedUnix(input.UpdatedUnix)
	return iuo
}

func (iu *InstructUpdate) SetInstruct(input *Instruct) *InstructUpdate {
	iu.SetCurrentUnix(input.CurrentUnix)
	iu.SetUpdatedUnix(input.UpdatedUnix)
	return iu
}

func (iu *InstructUpdate) SetInstructWithOptional(input *Instruct) *InstructUpdate {
	iu.SetCurrentUnix(input.CurrentUnix)
	iu.SetUpdatedUnix(input.UpdatedUnix)
	return iu
}

func (rc *ResourceCreate) SetResource(input *Resource) *ResourceCreate {
	rc.SetRid(input.Rid)
	rc.SetStatus(input.Status)
	rc.SetStep(input.Step)
	rc.SetRetried(input.Retried)
	rc.SetPriority(input.Priority)
	rc.SetRelate(input.Relate)
	rc.SetUpdatedUnix(input.UpdatedUnix)
	return rc
}

func (rc *ResourceCreate) SetResourceWithOptional(input *Resource) *ResourceCreate {
	rc.SetRid(input.Rid)
	rc.SetStatus(input.Status)
	rc.SetStep(input.Step)
	rc.SetRetried(input.Retried)
	rc.SetPriority(input.Priority)
	rc.SetRelate(input.Relate)
	rc.SetUpdatedUnix(input.UpdatedUnix)
	rc.SetComment(input.Comment)
	return rc
}

func (ruo *ResourceUpdateOne) SetResource(input *Resource) *ResourceUpdateOne {
	ruo.SetRid(input.Rid)
	ruo.SetStatus(input.Status)
	ruo.SetStep(input.Step)
	ruo.SetRetried(input.Retried)
	ruo.SetPriority(input.Priority)
	ruo.SetRelate(input.Relate)
	ruo.SetUpdatedUnix(input.UpdatedUnix)
	return ruo
}

func (ru *ResourceUpdate) SetResource(input *Resource) *ResourceUpdate {
	ru.SetRid(input.Rid)
	ru.SetStatus(input.Status)
	ru.SetStep(input.Step)
	ru.SetRetried(input.Retried)
	ru.SetPriority(input.Priority)
	ru.SetRelate(input.Relate)
	ru.SetUpdatedUnix(input.UpdatedUnix)
	return ru
}

func (ru *ResourceUpdate) SetResourceWithOptional(input *Resource) *ResourceUpdate {
	ru.SetRid(input.Rid)
	ru.SetStatus(input.Status)
	ru.SetStep(input.Step)
	ru.SetRetried(input.Retried)
	ru.SetPriority(input.Priority)
	ru.SetRelate(input.Relate)
	ru.SetUpdatedUnix(input.UpdatedUnix)
	ru.SetComment(input.Comment)
	return ru
}

func (vc *VersionCreate) SetVersion(input *Version) *VersionCreate {
	vc.SetCurrent(input.Current)
	vc.SetLast(input.Last)
	return vc
}

func (vc *VersionCreate) SetVersionWithOptional(input *Version) *VersionCreate {
	vc.SetCurrent(input.Current)
	vc.SetLast(input.Last)
	return vc
}

func (vuo *VersionUpdateOne) SetVersion(input *Version) *VersionUpdateOne {
	vuo.SetCurrent(input.Current)
	vuo.SetLast(input.Last)
	return vuo
}

func (vu *VersionUpdate) SetVersion(input *Version) *VersionUpdate {
	vu.SetCurrent(input.Current)
	vu.SetLast(input.Last)
	return vu
}

func (vu *VersionUpdate) SetVersionWithOptional(input *Version) *VersionUpdate {
	vu.SetCurrent(input.Current)
	vu.SetLast(input.Last)
	return vu
}
