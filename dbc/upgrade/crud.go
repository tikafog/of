// Code generated by ent, DO NOT EDIT.

package upgrade

import (
	"time"

	"github.com/tikafog/of/dbc/upgrade/update"
)

func (uc *UpdateCreate) SetUpdate(input *Update) *UpdateCreate {
	uc.SetCreatedUnix(input.CreatedUnix)
	uc.SetUpdatedUnix(input.UpdatedUnix)
	uc.SetOs(input.Os)
	uc.SetArch(input.Arch)
	uc.SetVersion(input.Version)
	uc.SetRid(input.Rid)
	uc.SetCrc32(input.Crc32)
	uc.SetAttr(input.Attr)
	uc.SetForcibly(input.Forcibly)
	uc.SetTruncate(input.Truncate)
	uc.SetTitle(input.Title)
	uc.SetDetail(input.Detail)
	return uc
}

func (uc *UpdateCreate) SetUpdateWithOptional(input *Update) *UpdateCreate {
	uc.SetCreatedUnix(input.CreatedUnix)
	uc.SetUpdatedUnix(input.UpdatedUnix)
	uc.SetOs(input.Os)
	uc.SetArch(input.Arch)
	uc.SetVersion(input.Version)
	uc.SetRid(input.Rid)
	uc.SetCrc32(input.Crc32)
	uc.SetAttr(input.Attr)
	uc.SetForcibly(input.Forcibly)
	uc.SetTruncate(input.Truncate)
	uc.SetTitle(input.Title)
	uc.SetDetail(input.Detail)
	uc.SetSign(input.Sign)
	return uc
}
func (uc *UpdateCreate) SoftDelete() *UpdateCreate {
	uc.SetDeletedUnix(time.Now().Unix())
	return uc
}

func (uuo *UpdateUpdateOne) SoftDelete() *UpdateUpdateOne {
	uuo.SetDeletedUnix(time.Now().Unix())
	return uuo
}

func (uuo *UpdateUpdateOne) SetUpdate(input *Update) *UpdateUpdateOne {
	uuo.SetUpdatedUnix(input.UpdatedUnix)
	uuo.SetOs(input.Os)
	uuo.SetArch(input.Arch)
	uuo.SetVersion(input.Version)
	uuo.SetRid(input.Rid)
	uuo.SetCrc32(input.Crc32)
	uuo.SetAttr(input.Attr)
	uuo.SetForcibly(input.Forcibly)
	uuo.SetTruncate(input.Truncate)
	uuo.SetTitle(input.Title)
	uuo.SetDetail(input.Detail)
	return uuo
}

func (uu *UpdateUpdate) SoftDelete() *UpdateUpdate {
	uu.SetDeletedUnix(time.Now().Unix())
	return uu
}

func (uu *UpdateUpdate) SetUpdate(input *Update) *UpdateUpdate {
	uu.SetUpdatedUnix(input.UpdatedUnix)
	uu.SetOs(input.Os)
	uu.SetArch(input.Arch)
	uu.SetVersion(input.Version)
	uu.SetRid(input.Rid)
	uu.SetCrc32(input.Crc32)
	uu.SetAttr(input.Attr)
	uu.SetForcibly(input.Forcibly)
	uu.SetTruncate(input.Truncate)
	uu.SetTitle(input.Title)
	uu.SetDetail(input.Detail)
	return uu
}

func (uu *UpdateUpdate) SetUpdateWithOptional(input *Update) *UpdateUpdate {
	uu.SetUpdatedUnix(input.UpdatedUnix)
	uu.SetOs(input.Os)
	uu.SetArch(input.Arch)
	uu.SetVersion(input.Version)
	uu.SetRid(input.Rid)
	uu.SetCrc32(input.Crc32)
	uu.SetAttr(input.Attr)
	uu.SetForcibly(input.Forcibly)
	uu.SetTruncate(input.Truncate)
	uu.SetTitle(input.Title)
	uu.SetDetail(input.Detail)
	uu.SetSign(input.Sign)
	return uu
}

func (uq *UpdateQuery) UseSoftDelete() *UpdateQuery {
	return uq.Where(update.DeletedUnixIsNil())
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
