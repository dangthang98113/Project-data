package Tuan1;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Component;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.GridLayout;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.EventObject;

import javax.swing.BorderFactory;
import javax.swing.Box;
import javax.swing.BoxLayout;
import javax.swing.JButton;
import javax.swing.JFormattedTextField;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JOptionPane;
import javax.swing.JPanel;
import javax.swing.JRadioButton;
import javax.swing.JTextField;
import javax.swing.border.Border;
import javax.swing.border.TitledBorder;

public class Bai3 extends JFrame implements ActionListener{

	JLabel lbTitle;
	JButton btnGiai, btnXoa, btnThoat;
	JTextField txtA,txtB,txtkq;
	JRadioButton radCong, radTru, radNhan, radChia;
	public Bai3(){
		super("CONG TRU NHA CHIA");
		setSize(4010,300);
		setVisible(true);
		setLocationRelativeTo(null);
		addControl();
		setResizable(false);
	}
	
	private void addControl() {
		// TODO Auto-generated method stub
		JPanel pnBorder = new JPanel();
		pnBorder.setLayout(new BorderLayout());
		JPanel pNorth = new JPanel();
		pNorth.add(lbTitle = new JLabel("CONG TRU NHAN CHIA"));
		Font fp = new Font("Aria", Font.BOLD,25);
		lbTitle.setFont(fp);
		pNorth.setBackground(Color.GREEN);
		pnBorder.add(pNorth,BorderLayout.NORTH);
		
		JPanel pnWest = new JPanel();
		pnWest.setLayout(new BoxLayout(pnWest, BoxLayout.Y_AXIS));
		btnGiai = new JButton("Giai");
		btnXoa = new JButton("Xoa");
		btnThoat = new JButton("Thoat");
		pnWest.add(btnGiai);
		pnWest(Box.createVerticalStrut(10));
		pnWest.add(btnXoa);
		pnWest(Box.createVerticalStrut(10));
		pnWest.add(btnThoat);
		pnBorder.add(pnWest,BorderLayout.WEST);
		pnWest.setBackground(Color.LIGHT_GRAY);;
		Border southBorder = BorderFactory.createLineBorder(Color.RED);
		TitledBorder southTitleBorder = new TitledBorder(southBorder,"Chon tac vu");
		pnWest.setBorder(southTitleBorder);
		
		JPanel pnSouth = new JPanel();
		pnSouth.setPreferredSize(new Dimension(0,30));
		pnSouth.setBackground(Color.PINK);
		JPanel pns1= new JPanel();
		pns1.setBackground(Color.BLUE);
		pnSouth.add(pns1);
		JPanel pns2= new JPanel();
		pns1.setBackground(Color.RED);
		pnSouth.add(pns2);
		JPanel pns3= new JPanel();
		pns1.setBackground(Color.YELLOW);
		pnSouth.add(pns3);
		pnBorder.add(pnSouth, BorderLayout.SOUTH);
		add(pnBorder);
		
		JPanel pnCenter = new JPanel();
		pnCenter.setLayout(new BoxLayout(pnCenter, BoxLayout.Y_AXIS));
		pnBorder.add(pnCenter,BorderLayout.CENTER);
		Border centerborder = BorderFactory.createLineBorder(Color.RED);
		TitledBorder centerTitleBorder = new TitledBorder(centerborder,"Nhap 2 so a vs b");
		pnCenter.setBorder(centerTitleBorder);
		JPanel pna = new JPanel();
		JLabel lbla = new JLabel("Nhap a");
		txtA = new JTextField(15);
		pna.add(lbla);
		pna.add(txtA);
		pnCenter.add(pna);
		JPanel pnb = new JPanel();
		JLabel lblb = new JLabel("Nhap b");
		txtB = new JTextField(15);
		pna.add(lblb);
		pna.add(txtB);
		pnCenter.add(pnb);
		JPanel pnc = new JPanel();
		JPanel pnpheptoan = new JPanel();
		pnpheptoan.setLayout(new GridLayout(2,2));
		pnpheptoan.setBorder(new TitledBorder(BorderFactory.createLineBorder(Color.BLACK),"Chon phep Toan"));
		radCong=new JRadioButton("Cong");
		radCong.setSelected(true);
		pnpheptoan.add(radCong);
		radNhan=new JRadioButton("Nhan");
		pnpheptoan.add(radNhan);
		radChia=new JRadioButton("Chia");
		pnpheptoan.add(radChia);
		pnc.add(pnpheptoan);
		pnCenter.add(pnc);
		
		JPanel pnkq = new JPanel();
		JLabel lblkq = new JLabel("ket qua");
		txtkq = new JTextField(15);
		pnkq.add(lblkq);
		pnkq.add(txtkq);
		pnCenter.add(pnkq);
		lbla.setPreferredSize(lblkq.getPreferredSize());
		lblb.setPreferredSize(lblkq.getPreferredSize());
		
		btnGiai.addActionListener(this);
		btnXoa.addActionListener(this);
	}
	private void pnWest(Component createVerticalStrut) {
		// TODO Auto-generated method stub
		
	}

	@Override
	public void actionPerformed(ActionEvent arg0) {
		// TODO Auto-generated method stub
		int a=0, b=0;
		String s1, s2;
		s1=txtA.getText();
		s2=txtB.getText();
		if(s1.equals("")){
			JOptionPane.showMessageDialog(this, "Ban chua nha a !!! ");
			txtA.requestFocus();
		}
		else
			if(s2.equals("")){
				JOptionPane.showMessageDialog(this, "Ban chua nha b !!! ");
				txtB.requestFocus();
			}
			else
				try{
					a= Integer.parseInt(txtA.getText());
				}
				catch(Exception ex){
					JOptionPane.showMessageDialog(this, "Nhap sai dinh dang !!!");
					txtA.selectAll();
					txtA.requestFocus();
				}
				try{
					b= Integer.parseInt(txtB.getText());
				}
				catch(Exception ex){
					JOptionPane.showMessageDialog(this, "Nhap sai dinh dang !!!");
					txtB.selectAll();
					txtB.requestFocus();
				}
				EventObject e = null;
				Object obj = e.getSource();
				if(obj.equals(btnGiai)){
					if(radCong.isSelected())
						txtkq.setText(Integer.toString(a+b));
					else
						if(radTru.isSelected())
							txtkq.setText(Integer.toString(a-b));
						else
							if(radNhan.isSelected())
								txtkq.setText(Integer.toString(a*b));
							else
								txtkq.setText(Double.toString(1.0*a/b));
				}
				else
					if(obj.equals(btnXoa)){
						txtA.setText("");
						txtB.setText("");
						txtkq.setText("");
						txtA.requestFocus();
					}
	}

}
