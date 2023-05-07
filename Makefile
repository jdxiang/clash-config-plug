include $(TOPDIR)/rules.mk

PKG_NAME:=clash-config-plug
PKG_VERSION:=0.0.1
PKG_RELEASE:=1

PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)-$(PKG_VERSION)
PKG_SOURCE:=$(PKG_BUILD_DIR)

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_USE_MIPS16:=0

GO_PKG:=clash-config-plug
GO_PKG_LDFLAGS:=-s -w
GO_PKG_LDFLAGS_X:= \
	$(GO_PKG)/version.Version=$(PKG_VERSION)

export GO111MODULE=on
export GOPROXY=https://mirrors.aliyun.com/goproxy/

include $(INCLUDE_DIR)/package.mk
include $(TOPDIR)/feeds/packages/lang/golang/golang-package.mk

define Package/$(PKG_NAME)
	SECTION:=net
	CATEGORY:=Network
	TITLE:=Open Clash cofig proxy plugin
	DEPENDS:=$(GO_ARCH_DEPENDS)
	URL:=https://www.baidu.com
	SUBMENU:=clash-config-plug
endef

define Package/$(PKG_NAME)/description
Open Clash cofig proxy plugin
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
	$(CP) ./src/* $(PKG_BUILD_DIR)/
endef

define Build/Configure

endef

define Build/Compile
	$(eval GO_PKG_BUILD_PKG:=$(GO_PKG))
	$(call GoPackage/Build/Configure)
	$(call GoPackage/Build/Compile)
	$(STAGING_DIR_HOST)/bin/upx --lzma --best $(GO_PKG_BUILD_BIN_DIR)/clash-config-plug
	chmod +wx $(GO_PKG_BUILD_BIN_DIR)/clash-config-plug
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(GO_PKG_BUILD_BIN_DIR)/clash-config-plug $(1)/usr/bin/clash-config-plug
	$(INSTALL_DIR) $(1)/usr/share/clash-config-plug
	$(CP) ./conf/* $(1)/usr/share/clash-config-plug/

endef
$(eval $(call GoBinPackage,$(PKG_NAME)))
$(eval $(call BuildPackage,$(PKG_NAME)))