using FreeERP.View;
namespace FreeERP
{
    partial class MainPage
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;
        public int mPageWidth, mPageHeight;


        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.tabControl1 = new System.Windows.Forms.TabControl();
            this.tabPage1 = new System.Windows.Forms.TabPage();
            this.soldManagerForm = new FreeERP.View.SoldManagerForm();
            this.tabPage2 = new System.Windows.Forms.TabPage();
            this.repoManagerForm = new FreeERP.View.RepoManagerForm();
            this.tabControl1.SuspendLayout();
            this.tabPage1.SuspendLayout();
            this.tabPage2.SuspendLayout();
            this.SuspendLayout(); this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.FixedSingle; // 禁止窗口大小调整
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen; // 窗口居中
            // 
            // tabControl1
            // 
            this.tabControl1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tabControl1.Location = new System.Drawing.Point(0, 0);
            this.tabControl1.Name = "tabControl1";
            this.tabControl1.SelectedIndex = 0;
            this.tabControl1.Size = new System.Drawing.Size(749, 346);
            this.tabControl1.TabIndex = 0;


            this.tabControl1.Controls.Add(this.tabPage1);
            this.tabControl1.Controls.Add(this.tabPage2);
            // 
            // tabPage1
            // 
            this.tabPage1.Name = "tabPage1";
            this.tabPage1.TabIndex = 0;
            this.tabPage1.Text = "销售管理";
            // 
            // soldManagerForm
            // 
            this.soldManagerForm.Dock = System.Windows.Forms.DockStyle.Fill;
            this.soldManagerForm.FormBorderStyle = System.Windows.Forms.FormBorderStyle.None;
            this.soldManagerForm.TopLevel = false;
            //this.soldManagerForm.WindowState = System.Windows.Forms.FormWindowState.Maximized;
            this.soldManagerForm.Name = "soldManagerForm";
            this.soldManagerForm.Text = "销售管理";
            this.soldManagerForm.Visible = false;
            this.tabPage1.Controls.Add(this.soldManagerForm);
            this.soldManagerForm.Show();
            // 
            // repoManagerForm
            // 
            this.tabPage2.Name = "tabPage2";
            this.tabPage2.TabIndex = 1;
            this.tabPage2.Text = "库存管理";
            this.repoManagerForm.Dock = System.Windows.Forms.DockStyle.Fill;
            this.repoManagerForm.FormBorderStyle = System.Windows.Forms.FormBorderStyle.None;
            this.repoManagerForm.TopLevel = false;
            //this.repoManagerForm.WindowState = System.Windows.Forms.FormWindowState.Maximized;
            this.repoManagerForm.Name = "repoManagerForm";
            this.repoManagerForm.Text = this.tabPage2.Text;
            this.repoManagerForm.Visible = false;
            this.tabPage2.Controls.Add(this.repoManagerForm);
            this.repoManagerForm.Show();


            // 
            // MainPage
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(8F, 15F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(749, 346);
            this.Controls.Add(this.tabControl1);
            this.Font = new System.Drawing.Font("Microsoft Sans Serif", 9F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.MaximizeBox = false;
            this.Name = "MainPage";
            this.Text = "欢迎使用FreeERP系统";
            this.tabControl1.ResumeLayout(false);
            this.tabPage1.ResumeLayout(false);
            this.tabPage2.ResumeLayout(false);
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TabControl tabControl1;
        private System.Windows.Forms.TabPage tabPage1;
        private System.Windows.Forms.TabPage tabPage2;
        private SoldManagerForm soldManagerForm;
        private RepoManagerForm repoManagerForm;

        // 销售管理 // 库存管理
    }
}

