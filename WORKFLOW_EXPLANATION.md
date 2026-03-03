# GitHub Actions Workflow - Detailed Explanation

## 1. Source of Truth for App Version

### Current Implementation (After Fix)

**Docker Image Tags:**
- Source: Git tags (e.g., `v1.0.0`) via `docker/metadata-action`
- Extracted from: `github.ref_name` when `github.ref_type == "tag"`
- Examples: `v1.0.0`, `v1.0`, `v1`, `latest`

**Go Binary Version (`main.Version`):**
- Source: Build argument `VERSION` passed to Dockerfile
- Extracted from: Same git tag/ref as Docker tags
- Passed via: `-ldflags "-X main.Version=${VERSION}"` in Dockerfile
- Examples: `1.0.0` (from tag `v1.0.0`), `dev-abc123` (from main branch)

### Version Synchronization

✅ **NOW TIED TOGETHER**: The workflow extracts the version from the git ref/tag and passes it to both:
1. Docker metadata action (for image tags)
2. Docker build args (for Go binary version)

**Version Extraction Logic:**
- **Tag event** (e.g., `v1.0.0`): Extracts `1.0.0` (removes 'v' prefix)
- **Main/master branch**: Uses `dev-<commit-sha>` as version
- **Other branches/PRs**: Uses `dev-<branch>-<commit-sha>`

**Result**: The Docker image tag and the Go binary version will always match:
- Image `kaungmyathan/devops-bse-bootcamp:v1.0.0` contains binary with `Version="1.0.0"`
- Image `kaungmyathan/devops-bse-bootcamp:latest` (from main) contains binary with `Version="dev-<sha>"`

---

## 2. Failure Behavior

### Scenario: Tests Pass ✅, Docker Build Fails ❌

**Step-by-Step Execution:**

1. **Checkout code** ✅
   - Status: Success
   - Action: Continues

2. **Set up Go** ✅
   - Status: Success
   - Action: Continues

3. **Run Go tests** ✅
   - Status: Success (all tests pass)
   - Action: Continues

4. **Set up Docker Buildx** ✅
   - Status: Success
   - Action: Continues

5. **Log in to Docker Hub** ✅ (if not PR)
   - Status: Success
   - Action: Continues

6. **Extract metadata** ✅
   - Status: Success
   - Action: Continues

7. **Extract version** ✅
   - Status: Success
   - Action: Continues

8. **Build and push Docker image** ❌
   - Status: **FAILURE**
   - What happens:
     - Docker build command fails (e.g., syntax error, dependency issue)
     - Build step exits with non-zero exit code
     - **No image is pushed** (build must succeed before push)
     - **No later steps execute** (workflow stops here)

**Workflow Status:**
- Overall workflow status: ❌ **FAILED**
- GitHub UI shows: Red X with "Build and push Docker image" step failed
- No Docker image is created or pushed
- Previous steps (tests, etc.) still show as ✅ in logs

**What This Means:**
- The workflow is **fail-fast**: If any step fails, the workflow stops
- Tests passing doesn't guarantee a successful build
- You must fix the Docker build issue and re-run the workflow
- No partial artifacts are left behind (clean failure)

**For a Junior Developer:**
> "Think of it like a production line. Each step must complete successfully before moving to the next. If the Docker build fails, it's like the assembly line stops - nothing gets packaged or shipped, even though all the quality checks (tests) passed. You need to fix the build issue and start the process again."

---

## 3. Multi-Platform Builds

### Current Implementation

**Platforms Supported:**
- `linux/amd64` (Intel/AMD 64-bit servers)
- `linux/arm64` (ARM 64-bit, e.g., M1/M2 Macs, ARM servers)

**Changes Made:**

1. **Workflow (`docker.yml`):**
   ```yaml
   platforms: linux/amd64,linux/arm64
   ```
   - Builds the same image for both architectures
   - Creates a multi-arch manifest
   - Docker automatically selects the correct image for the user's platform

2. **Dockerfile:**
   - No changes needed! The Dockerfile is already architecture-agnostic
   - Go's cross-compilation handles architecture differences
   - Buildx automatically builds for each specified platform

**How It Works:**
- Buildx creates separate images for each platform
- Combines them into a single multi-arch manifest
- When users pull the image, Docker automatically selects the correct architecture
- Example: M1 Mac pulls `linux/arm64`, Intel server pulls `linux/amd64`

**Benefits:**
- ✅ Works on both Intel/AMD and ARM-based systems
- ✅ Same image tag works everywhere
- ✅ No need for separate builds or tags

**Build Time:**
- Takes longer (builds 2 images instead of 1)
- But enables broader compatibility

---

## Summary

### Version Management
- ✅ Docker tags and Go binary versions are now synchronized
- ✅ Single source of truth: Git refs/tags
- ✅ No version drift possible

### Failure Handling
- ✅ Fail-fast behavior: Workflow stops on first failure
- ✅ Tests can pass but build can still fail
- ✅ No partial artifacts created

### Multi-Platform
- ✅ Supports both amd64 and arm64
- ✅ Automatic platform selection
- ✅ No Dockerfile changes needed
