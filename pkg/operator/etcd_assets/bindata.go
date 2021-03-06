// Code generated by go-bindata.
// sources:
// bindata/etcd/cm.yaml
// bindata/etcd/defaultconfig.yaml
// bindata/etcd/etcd-common-tools
// bindata/etcd/etcd-member-remove.sh
// bindata/etcd/etcd-snapshot-backup.sh
// bindata/etcd/etcd-snapshot-restore.sh
// bindata/etcd/ns.yaml
// bindata/etcd/pod-cm.yaml
// bindata/etcd/pod.yaml
// bindata/etcd/restore-pod-cm.yaml
// bindata/etcd/restore-pod.yaml
// bindata/etcd/sa.yaml
// bindata/etcd/scripts-cm.yaml
// bindata/etcd/svc.yaml
// DO NOT EDIT!

package etcd_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _etcdCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: config
data:
  config.yaml:
`)

func etcdCmYamlBytes() ([]byte, error) {
	return _etcdCmYaml, nil
}

func etcdCmYaml() (*asset, error) {
	bytes, err := etcdCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdDefaultconfigYaml = []byte(`apiVersion: kubecontrolplane.config.openshift.io/v1
kind: EtcdConfig
`)

func etcdDefaultconfigYamlBytes() ([]byte, error) {
	return _etcdDefaultconfigYaml, nil
}

func etcdDefaultconfigYaml() (*asset, error) {
	bytes, err := etcdDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdEtcdCommonTools = []byte(`# Common environment variables
ASSET_DIR="/home/core/assets"
CONFIG_FILE_DIR="/etc/kubernetes"
MANIFEST_DIR="${CONFIG_FILE_DIR}/manifests"
ETCD_DATA_DIR="/var/lib/etcd"
ETCD_DATA_DIR_BACKUP="/var/lib/etcd-backup"
MANIFEST_STOPPED_DIR="${ASSET_DIR}/manifests-stopped"
RESTORE_ETCD_POD_YAML="${CONFIG_FILE_DIR}/static-pod-resources/etcd-certs/configmaps/restore-etcd-pod/pod.yaml"
ETCDCTL_BIN_DIR="${CONFIG_FILE_DIR}/static-pod-resources/bin"
PATH=${PATH}:${ETCDCTL_BIN_DIR}

# download etcdctl from upstream release assets
function dl_etcdctl {
  local etcdimg=${ETCD_IMAGE}
  local etcdctr=$(podman create ${etcdimg})
  local etcdmnt=$(podman mount "${etcdctr}")
  [ ! -d ${ETCDCTL_BIN_DIR} ] && mkdir -p ${ETCDCTL_BIN_DIR}
  cp ${etcdmnt}/bin/etcdctl ${ETCDCTL_BIN_DIR}/
  umount "${etcdmnt}"
  podman rm "${etcdctr}"
  etcdctl version
}
`)

func etcdEtcdCommonToolsBytes() ([]byte, error) {
	return _etcdEtcdCommonTools, nil
}

func etcdEtcdCommonTools() (*asset, error) {
	bytes, err := etcdEtcdCommonToolsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/etcd-common-tools", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdEtcdMemberRemoveSh = []byte(`#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o errtrace

# example
# sudo ./etcd-member-remove.sh $etcd_name

if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root"
  exit 1
fi

function usage {
    echo 'The name of the etcd member to remove is required: ./etcd-member-remove.sh $etcd_name'
    exit 1
}

### main
if [ "$1" == "" ]; then
    usage
fi

NAME="$1"

source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd.env
source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd-common-tools

# Download etcdctl binary
dl_etcdctl

# If the 1st field or the 3rd field of the member list exactly matches with the name, then get its ID. Note 3rd field has extra space to match.
ID=$(etcdctl member list | awk -F,  "\$1 ~ /^${NAME}$/ || \$3 ~ /^\s${NAME}$/ { print \$1 }")
if [ "$?" -ne 0 ] || [ -z "$ID" ]; then
    echo "could not find etcd member $NAME to remove."
    exit 1
fi

# Remove the member using ID
etcdctl member remove $ID
`)

func etcdEtcdMemberRemoveShBytes() ([]byte, error) {
	return _etcdEtcdMemberRemoveSh, nil
}

func etcdEtcdMemberRemoveSh() (*asset, error) {
	bytes, err := etcdEtcdMemberRemoveShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/etcd-member-remove.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdEtcdSnapshotBackupSh = []byte(`#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o errtrace

# example
# etcd-snapshot-backup.sh $path-to-snapshot

if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root"
  exit 1
fi

function usage {
    echo 'Path to backup dir required: ./etcd-snapshot-backup.sh <path-to-backup-dir>'
    exit 1
}

#backup latest static pod resources for kube-apiserver
function backup_latest_kube_static_resources {
  echo "Trying to backup latest static pod resources.."
  LATEST_STATIC_POD_DIR=$(ls -vd "${CONFIG_FILE_DIR}"/static-pod-resources/kube-apiserver-pod-[0-9]* | tail -1) || true
  if [ -z "$LATEST_STATIC_POD_DIR" ]; then
      echo "error finding static-pod-resources"
      exit 1
  fi

  LATEST_ETCD_STATIC_POD_DIR=$(ls -vd "${CONFIG_FILE_DIR}"/static-pod-resources/etcd-pod-[0-9]* | tail -1) || true
  if [ -z "$LATEST_ETCD_STATIC_POD_DIR" ]; then
      echo "error finding static-pod-resources"
      exit 1
  fi

  # tar up the static kube resources, with the path relative to CONFIG_FILE_DIR
  tar -cpzf $BACKUP_TAR_FILE -C ${CONFIG_FILE_DIR} ${LATEST_STATIC_POD_DIR#$CONFIG_FILE_DIR/} ${LATEST_ETCD_STATIC_POD_DIR#$CONFIG_FILE_DIR/}
}


# main
# If the first argument is missing, or it is an existing file, then print usage and exit
if [ -z "$1" ] || [ -f "$1" ]; then
  usage
fi

if [ ! -d "$1" ]; then
  mkdir -p $1
fi

BACKUP_DIR="$1"
DATESTRING=$(date "+%F_%H%M%S")
BACKUP_TAR_FILE=${BACKUP_DIR}/static_kuberesources_${DATESTRING}.tar.gz
SNAPSHOT_FILE="${BACKUP_DIR}/snapshot_${DATESTRING}.db"

trap "rm -f ${BACKUP_TAR_FILE} ${SNAPSHOT_FILE}" ERR

source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd.env
source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd-common-tools

dl_etcdctl
backup_latest_kube_static_resources
etcdctl snapshot save ${SNAPSHOT_FILE}
echo "snapshot db and kube resources are successfully saved to ${BACKUP_DIR}!"
`)

func etcdEtcdSnapshotBackupShBytes() ([]byte, error) {
	return _etcdEtcdSnapshotBackupSh, nil
}

func etcdEtcdSnapshotBackupSh() (*asset, error) {
	bytes, err := etcdEtcdSnapshotBackupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/etcd-snapshot-backup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdEtcdSnapshotRestoreSh = []byte(`#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o errtrace

# example
# ./etcd-snapshot-restore.sh $path-to-backup 

if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root"
  exit 1
fi

source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd.env
source /etc/kubernetes/static-pod-resources/etcd-certs/configmaps/etcd-scripts/etcd-common-tools

function usage {
  echo 'Path to the directory containing backup files is required: ./etcd-snapshot-restore.sh <path-to-backup>'
  echo 'The backup directory is expected to be contain two files:'
  echo '        1. etcd snapshot'
  echo '        2. A copy of the Static POD resources at the time of backup'
  exit 1
}

# If the argument is not passed, or if it is not a directory, print usage and exit.
if [ "$1" == "" ] || [ ! -d "$1" ]; then
  usage
fi

BACKUP_DIR="$1"
BACKUP_FILE=$(ls -vd "${BACKUP_DIR}"/static_kuberesources*.tar.gz | tail -1) || true
SNAPSHOT_FILE=$(ls -vd "${BACKUP_DIR}"/snapshot*.db | tail -1) || true

if [ ! -f "${SNAPSHOT_FILE}" ]; then
  echo "etcd snapshot ${SNAPSHOT_FILE} does not exist."
  exit 1
fi

# Move manifests and stop static pods
if [ ! -d "$MANIFEST_STOPPED_DIR" ]; then
  mkdir $MANIFEST_STOPPED_DIR
fi

# Move static pod manifests out of MANIFEST_DIR
find ${MANIFEST_DIR} \
  -maxdepth 1 \
  -type f \
  -printf '...stopping %P\n' \
  -exec mv {} ${MANIFEST_STOPPED_DIR} \;

# Wait for pods to stop
sleep 30

# //TO DO: verify using crictl that etcd and other pods stopped.

# Remove data dir
echo "Moving etcd data-dir ${ETCD_DATA_DIR}/member to ${ETCD_DATA_DIR_BACKUP}"
[ ! -d ${ETCD_DATA_DIR_BACKUP} ]  && mkdir -p ${ETCD_DATA_DIR_BACKUP}
mv ${ETCD_DATA_DIR}/member ${ETCD_DATA_DIR_BACKUP}/member

# Copy snapshot to backupdir
if [ ! -d ${ETCD_DATA_DIR_BACKUP} ]; then
  mkdir -p ${ETCD_DATA_DIR_BACKUP}
fi
cp -p ${SNAPSHOT_FILE} ${ETCD_DATA_DIR_BACKUP}/snapshot.db

# Copy etcd restore pod yaml
cp -p ${RESTORE_ETCD_POD_YAML} ${MANIFEST_DIR}/etcd-pod.yaml

# Restore static pod resources
tar -C ${CONFIG_FILE_DIR} -xzf ${BACKUP_FILE} static-pod-resources
`)

func etcdEtcdSnapshotRestoreShBytes() ([]byte, error) {
	return _etcdEtcdSnapshotRestoreSh, nil
}

func etcdEtcdSnapshotRestoreSh() (*asset, error) {
	bytes, err := etcdEtcdSnapshotRestoreShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/etcd-snapshot-restore.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-etcd
  labels:
    openshift.io/run-level: "0"
`)

func etcdNsYamlBytes() ([]byte, error) {
	return _etcdNsYaml, nil
}

func etcdNsYaml() (*asset, error) {
	bytes, err := etcdNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdPodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: etcd-pod
data:
  pod.yaml:
  forceRedeploymentReason:
  version:
`)

func etcdPodCmYamlBytes() ([]byte, error) {
	return _etcdPodCmYaml, nil
}

func etcdPodCmYaml() (*asset, error) {
	bytes, err := etcdPodCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: etcd
  namespace: openshift-etcd
  labels:
    app: etcd
    k8s-app: etcd
    etcd: "true"
    revision: "REVISION"
spec:
  initContainers:
    - name: etcd-quorum-guard-copy
      image: ${IMAGE}
      imagePullPolicy: IfNotPresent
      terminationMessagePolicy: FallbackToLogsOnError
      command:
        - /bin/sh
        - -c
        - |
          #!/bin/sh
          set -euo pipefail

          cp /etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.crt /etc/kubernetes/etcd-backup-dir/system:etcd-peer-NODE_NAME.crt
          cp /etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.key /etc/kubernetes/etcd-backup-dir/system:etcd-peer-NODE_NAME.key
      resources:
        requests:
          memory: 60Mi
          cpu: 30m
      securityContext:
        privileged: true
      volumeMounts:
        - mountPath: /etc/kubernetes/etcd-backup-dir
          name: etcd-backup-dir
        - mountPath: /etc/kubernetes/static-pod-resources
          name: resource-dir
        - mountPath: /etc/kubernetes/static-pod-certs
          name: cert-dir
  containers:
  # The etcdctl container should always be first. It is intended to be used
  # to open a remote shell via ` + "`" + `oc rsh` + "`" + ` that is ready to run ` + "`" + `etcdctl` + "`" + `.
  - name: etcdctl
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command:
      - "/bin/bash"
      - "-c"
      - "trap TERM INT; sleep infinity & wait"
    resources:
      requests:
        memory: 60Mi
        cpu: 30m
    volumeMounts:
      - mountPath: /etc/kubernetes/manifests
        name: static-pod-dir
      - mountPath: /etc/kubernetes/etcd-backup-dir
        name: etcd-backup-dir
      - mountPath: /etc/kubernetes/static-pod-resources
        name: resource-dir
      - mountPath: /etc/kubernetes/static-pod-certs
        name: cert-dir
      - mountPath: /var/lib/etcd/
        name: data-dir
    env:
${COMPUTED_ENV_VARS}
  - name: etcd
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command:
      - /bin/sh
      - -c
      - |
        #!/bin/sh
        set -euo pipefail

        etcdctl member list || true

        # this has a non-zero return code if the command is non-zero.  If you use an export first, it doesn't and you
        # will succeed when you should fail.
        ETCD_INITIAL_CLUSTER=$(discover-etcd-initial-cluster \
          --cacert=/etc/kubernetes/static-pod-certs/configmaps/etcd-serving-ca/ca-bundle.crt \
          --cert=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.crt \
          --key=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.key \
          --endpoints=${ALL_ETCD_ENDPOINTS} \
          --data-dir=/var/lib/etcd \
          --target-peer-url-host=${NODE_NODE_ENVVAR_NAME_ETCD_URL_HOST} \
          --target-name=NODE_NAME)
         export ETCD_INITIAL_CLUSTER

        # at this point we know this member is added.  To support a transition, we must remove the old etcd pod.
        # move it somewhere safe so we can retrieve it again later if something goes badly.
        mv /etc/kubernetes/manifests/etcd-member.yaml /etc/kubernetes/etcd-backup-dir || true

        # we cannot use the "normal" port conflict initcontainer because when we upgrade, the existing static pod will never yield,
        # so we do the detection in etcd container itsefl.
        echo -n "Waiting for ports 2379, 2380 and 9978 to be released."
        while [ -n "$(lsof -ni :2379)$(lsof -ni :2380)$(lsof -ni :9978)" ]; do
          echo -n "."
          sleep 1
        done

        export ETCD_NAME=${NODE_NODE_ENVVAR_NAME_ETCD_NAME}
        env | grep ETCD | grep -v NODE

        set -x
        exec etcd \
          --initial-advertise-peer-urls=https://${NODE_NODE_ENVVAR_NAME_IP}:2380 \
          --cert-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-serving/etcd-serving-NODE_NAME.crt \
          --key-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-serving/etcd-serving-NODE_NAME.key \
          --trusted-ca-file=/etc/kubernetes/static-pod-certs/configmaps/etcd-serving-ca/ca-bundle.crt \
          --client-cert-auth=true \
          --peer-cert-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.crt \
          --peer-key-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.key \
          --peer-trusted-ca-file=/etc/kubernetes/static-pod-certs/configmaps/etcd-peer-client-ca/ca-bundle.crt \
          --peer-client-cert-auth=true \
          --advertise-client-urls=https://${NODE_NODE_ENVVAR_NAME_IP}:2379 \
          --listen-client-urls=https://${LISTEN_ON_ALL_IPS}:2379 \
          --listen-peer-urls=https://${LISTEN_ON_ALL_IPS}:2380 \
          --listen-metrics-urls=https://${LISTEN_ON_ALL_IPS}:9978 ||  mv /etc/kubernetes/etcd-backup-dir/etcd-member.yaml /etc/kubernetes/manifests
    env:
${COMPUTED_ENV_VARS}
    resources:
      requests:
        memory: 600Mi
        cpu: 300m
    readinessProbe:
      exec:
        command:
          - /bin/sh
          - -ec
          - "lsof -n -i :2380 | grep LISTEN"
      failureThreshold: 3
      initialDelaySeconds: 3
      periodSeconds: 5
      successThreshold: 1
      timeoutSeconds: 5
    securityContext:
      privileged: true
    volumeMounts:
      - mountPath: /etc/kubernetes/manifests
        name: static-pod-dir
      - mountPath: /etc/kubernetes/etcd-backup-dir
        name: etcd-backup-dir
      - mountPath: /etc/kubernetes/static-pod-resources
        name: resource-dir
      - mountPath: /etc/kubernetes/static-pod-certs
        name: cert-dir
      - mountPath: /var/lib/etcd/
        name: data-dir
  - name: etcd-metrics
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command:
      - /bin/sh
      - -c
      - |
        #!/bin/sh
        set -euo pipefail

        export ETCD_NAME=${NODE_NODE_ENVVAR_NAME_ETCD_NAME}

        exec etcd grpc-proxy start \
          --endpoints https://${NODE_NODE_ENVVAR_NAME_ETCD_URL_HOST}:9978 \
          --metrics-addr https://${LISTEN_ON_ALL_IPS}:9979 \
          --listen-addr ${LOCALHOST_IP}:9977 \
          --key /etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.key \
          --key-file /etc/kubernetes/static-pod-certs/secrets/etcd-all-serving-metrics/etcd-serving-metrics-NODE_NAME.key \
          --cert /etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.crt \
          --cert-file /etc/kubernetes/static-pod-certs/secrets/etcd-all-serving-metrics/etcd-serving-metrics-NODE_NAME.crt \
          --cacert /etc/kubernetes/static-pod-certs/configmaps/etcd-peer-client-ca/ca-bundle.crt \
          --trusted-ca-file /etc/kubernetes/static-pod-certs/configmaps/etcd-metrics-proxy-serving-ca/ca-bundle.crt
    env:
${COMPUTED_ENV_VARS}
    resources:
      requests:
        memory: 200Mi
        cpu: 100m
    securityContext:
      privileged: true
    volumeMounts:
      - mountPath: /etc/kubernetes/static-pod-resources
        name: resource-dir
      - mountPath: /etc/kubernetes/static-pod-certs
        name: cert-dir
      - mountPath: /var/lib/etcd/
        name: data-dir
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
    - hostPath:
        path: /etc/kubernetes/manifests
      name: static-pod-dir
    - hostPath:
        path: /etc/kubernetes/static-pod-resources/etcd-member
      name: etcd-backup-dir
    - hostPath:
        path: /etc/kubernetes/static-pod-resources/etcd-pod-REVISION
      name: resource-dir
    - hostPath:
        path: /etc/kubernetes/static-pod-resources/etcd-certs
      name: cert-dir
    - hostPath:
        path: /var/lib/etcd
        type: ""
      name: data-dir

`)

func etcdPodYamlBytes() ([]byte, error) {
	return _etcdPodYaml, nil
}

func etcdPodYaml() (*asset, error) {
	bytes, err := etcdPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdRestorePodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: restore-etcd-pod
data:
  pod.yaml:
`)

func etcdRestorePodCmYamlBytes() ([]byte, error) {
	return _etcdRestorePodCmYaml, nil
}

func etcdRestorePodCmYaml() (*asset, error) {
	bytes, err := etcdRestorePodCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/restore-pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdRestorePodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: etcd
  namespace: openshift-etcd
  labels:
    app: etcd
    k8s-app: etcd
    etcd: "true"
    revision: "REVISION"
spec:
  containers:
  - name: etcd
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command:
      - /bin/sh
      - -c
      - |
        #!/bin/sh
        set -euo pipefail

        export ETCD_NAME=${NODE_NODE_ENVVAR_NAME_ETCD_NAME}
        export ETCD_INITIAL_CLUSTER="${ETCD_NAME}=https://${NODE_NODE_ENVVAR_NAME_ETCD_URL_HOST}:2380"
        env | grep ETCD | grep -v NODE
        export ETCD_NODE_PEER_URL=https://${NODE_NODE_ENVVAR_NAME_ETCD_URL_HOST}:2380

        # checking if data directory is empty, if not etcdctl restore will fail
        if [ ! -z $(ls -A "/var/lib/etcd") ]; then
          echo "please delete the contents of data directory before restoring, running the restore script will do this for you"
          exit 1
        fi

        # check if we have backup file to be restored
        # if the file exist, check if it has not changed size in last 5 seconds
        if [ ! -f /var/lib/etcd-backup/snapshot.db ]; then
          echo "please make a copy of the snapshot db file, then move that copy to /var/lib/etcd-backup/snapshot.db"
          exit 1
        else
          filesize=$(stat --format=%s "/var/lib/etcd-backup/snapshot.db")
          sleep 5
          newfilesize=$(stat --format=%s "/var/lib/etcd-backup/snapshot.db")
          if [ "$filesize" != "$newfilesize" ]; then
            echo "file size has changed since last 5 seconds, retry sometime after copying is complete"
            exit 1
          fi
        fi

        UUID=$(uuidgen)
        echo "restoring to a single node cluster"
        ETCDCTL_API=3 /usr/bin/etcdctl snapshot restore /var/lib/etcd-backup/snapshot.db \
         --name  $ETCD_NAME \
         --initial-cluster=$ETCD_INITIAL_CLUSTER \
         --initial-cluster-token "openshift-etcd-{$UUID:0:10}" \
         --initial-advertise-peer-urls $ETCD_NODE_PEER_URL \
         --data-dir="/var/lib/etcd/restore-{$UUID:0:10}"

        mv /var/lib/etcd/restore-{$UUID:0:10}/* /var/lib/etcd/

        set -x
        exec etcd \
          --initial-advertise-peer-urls=https://${NODE_NODE_ENVVAR_NAME_IP}:2380 \
          --cert-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-serving/etcd-serving-NODE_NAME.crt \
          --key-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-serving/etcd-serving-NODE_NAME.key \
          --trusted-ca-file=/etc/kubernetes/static-pod-certs/configmaps/etcd-serving-ca/ca-bundle.crt \
          --client-cert-auth=true \
          --peer-cert-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.crt \
          --peer-key-file=/etc/kubernetes/static-pod-certs/secrets/etcd-all-peer/etcd-peer-NODE_NAME.key \
          --peer-trusted-ca-file=/etc/kubernetes/static-pod-certs/configmaps/etcd-peer-client-ca/ca-bundle.crt \
          --peer-client-cert-auth=true \
          --advertise-client-urls=https://${NODE_NODE_ENVVAR_NAME_IP}:2379 \
          --listen-client-urls=https://${LISTEN_ON_ALL_IPS}:2379 \
          --listen-peer-urls=https://${LISTEN_ON_ALL_IPS}:2380 \
          --listen-metrics-urls=https://${LISTEN_ON_ALL_IPS}:9978 ||  mv /etc/kubernetes/etcd-backup-dir/etcd-member.yaml /etc/kubernetes/manifests
    env:
${COMPUTED_ENV_VARS}
    resources:
      requests:
        memory: 600Mi
        cpu: 300m
    readinessProbe:
      exec:
        command:
          - /bin/sh
          - -ec
          - "lsof -n -i :2380 | grep LISTEN"
      failureThreshold: 3
      initialDelaySeconds: 3
      periodSeconds: 5
      successThreshold: 1
      timeoutSeconds: 5
    securityContext:
      privileged: true
    volumeMounts:
      - mountPath: /etc/kubernetes/manifests
        name: static-pod-dir
      - mountPath: /etc/kubernetes/etcd-backup-dir
        name: etcd-backup-dir
      - mountPath: /etc/kubernetes/static-pod-certs
        name: cert-dir
      - mountPath: /var/lib/etcd/
        name: data-dir
      - mountPath: /var/lib/etcd-backup/
        name: backup-dir
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
    - hostPath:
        path: /etc/kubernetes/manifests
      name: static-pod-dir
    - hostPath:
        path: /etc/kubernetes/static-pod-resources/etcd-member
      name: etcd-backup-dir
    - hostPath:
        path: /etc/kubernetes/static-pod-resources/etcd-certs
      name: cert-dir
    - hostPath:
        path: /var/lib/etcd
        type: ""
      name: data-dir
    - hostPath:
        path: /var/lib/etcd-backup
        type: ""
      name: backup-dir

`)

func etcdRestorePodYamlBytes() ([]byte, error) {
	return _etcdRestorePodYaml, nil
}

func etcdRestorePodYaml() (*asset, error) {
	bytes, err := etcdRestorePodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/restore-pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-etcd
  name: etcd-sa
`)

func etcdSaYamlBytes() ([]byte, error) {
	return _etcdSaYaml, nil
}

func etcdSaYaml() (*asset, error) {
	bytes, err := etcdSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdScriptsCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: etcd-scripts
data:
  etcd.env:
  etcd-snapshot-restore.sh:
  etcd-snapshot-backup.sh:
  etcd-member-remove.sh:
  etcd-common-tools:
`)

func etcdScriptsCmYamlBytes() ([]byte, error) {
	return _etcdScriptsCmYaml, nil
}

func etcdScriptsCmYaml() (*asset, error) {
	bytes, err := etcdScriptsCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/scripts-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcdSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-etcd
  name: etcd
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    etcd: "true"
  ports:
  - name: https
    port: 443
    targetPort: 10257
`)

func etcdSvcYamlBytes() ([]byte, error) {
	return _etcdSvcYaml, nil
}

func etcdSvcYaml() (*asset, error) {
	bytes, err := etcdSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etcd/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"etcd/cm.yaml":                  etcdCmYaml,
	"etcd/defaultconfig.yaml":       etcdDefaultconfigYaml,
	"etcd/etcd-common-tools":        etcdEtcdCommonTools,
	"etcd/etcd-member-remove.sh":    etcdEtcdMemberRemoveSh,
	"etcd/etcd-snapshot-backup.sh":  etcdEtcdSnapshotBackupSh,
	"etcd/etcd-snapshot-restore.sh": etcdEtcdSnapshotRestoreSh,
	"etcd/ns.yaml":                  etcdNsYaml,
	"etcd/pod-cm.yaml":              etcdPodCmYaml,
	"etcd/pod.yaml":                 etcdPodYaml,
	"etcd/restore-pod-cm.yaml":      etcdRestorePodCmYaml,
	"etcd/restore-pod.yaml":         etcdRestorePodYaml,
	"etcd/sa.yaml":                  etcdSaYaml,
	"etcd/scripts-cm.yaml":          etcdScriptsCmYaml,
	"etcd/svc.yaml":                 etcdSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"etcd": {nil, map[string]*bintree{
		"cm.yaml":                  {etcdCmYaml, map[string]*bintree{}},
		"defaultconfig.yaml":       {etcdDefaultconfigYaml, map[string]*bintree{}},
		"etcd-common-tools":        {etcdEtcdCommonTools, map[string]*bintree{}},
		"etcd-member-remove.sh":    {etcdEtcdMemberRemoveSh, map[string]*bintree{}},
		"etcd-snapshot-backup.sh":  {etcdEtcdSnapshotBackupSh, map[string]*bintree{}},
		"etcd-snapshot-restore.sh": {etcdEtcdSnapshotRestoreSh, map[string]*bintree{}},
		"ns.yaml":                  {etcdNsYaml, map[string]*bintree{}},
		"pod-cm.yaml":              {etcdPodCmYaml, map[string]*bintree{}},
		"pod.yaml":                 {etcdPodYaml, map[string]*bintree{}},
		"restore-pod-cm.yaml":      {etcdRestorePodCmYaml, map[string]*bintree{}},
		"restore-pod.yaml":         {etcdRestorePodYaml, map[string]*bintree{}},
		"sa.yaml":                  {etcdSaYaml, map[string]*bintree{}},
		"scripts-cm.yaml":          {etcdScriptsCmYaml, map[string]*bintree{}},
		"svc.yaml":                 {etcdSvcYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
